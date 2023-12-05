package main

import (
	"database/sql"
	"fmt"
)

func main() {
	bb := txInitBb()
	//执行事务
	transactionDemo(bb)
}

/*
初始化数据库
*/
func txInitBb() (db *sql.DB) {
	dsn := "root:password@tcp(127.0.0.1:3306)/cygn?charset=utf8"
	//打开数据库连接
	db, err := sql.Open("mysql", dsn)
	//设置链接数
	db.SetMaxOpenConns(10)
	//设置连接池最大限制链接数
	db.SetMaxIdleConns(100)
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("尝试与数据库建立连接失败", err)
	}
	if err != nil {
		fmt.Printf("连接数据库失败", err)
	}

	return db
}

// 事务操作示例
func transactionDemo(bb *sql.DB) {
	tx, err := bb.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	fmt.Println("事务提交啦...")
	tx.Commit() // 提交事务

	fmt.Println("exec trans success!")
}
