package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	stop := make(chan struct{})
	g := new(errgroup.Group)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "Hello World!") })
	// http server
	g.Go(func() error {
		return server(":8080", mux, stop)
	})
	//signal
	g.Go(func() error {
		return signalProcess(stop)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Wait err===》%s", err)
	} else {
		fmt.Println("End")
	}
}

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func signalProcess(stop chan<- struct{}) error {
	c := make(chan os.Signal)
	signal.Notify(c)
	s := <-c
	close(stop)

	return errors.New(fmt.Sprintf("signal:%s shutdown", s))
}
