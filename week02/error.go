package main

import (
	"database/sql"
	"fmt"
)

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// if !r.rows.Next() {
//		if err := r.rows.Err(); err != nil {
//			return err
//		}
//		return ErrNoRows
//	}
// 通过查看Row.Scan方法是在记录为空时返回ErrNoRows是正常的错误返回，dao层应该自己处理掉这个error不需要抛给上层

type rows = map[string]string

func main() {
	row, err := dao()
	fmt.Println(row, err)
}

func dao() (rows, error) {
	var r rows
	err := scan(r)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return r, nil
}

func scan(dest ...interface{}) error {
	return sql.ErrNoRows
}
