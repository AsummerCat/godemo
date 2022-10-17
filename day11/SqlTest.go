package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/************数据库的基本使用*********/

/*
初始化数据库
*/
func initBb() (db *sql.DB) {
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

/*
查询单行
*/
func queryRowDemo(db *sql.DB) {
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow("select ACCOUNT_ID,DISPLAY_NAME from c_account limit 1").Scan(&u.accountId, &u.displayName)
	if err != nil {
		fmt.Printf("未找到结果", err)
	}
	fmt.Printf("单行查询结果:"+"id:%d name:%s age:%d\n", u.accountId, u.displayName)

}

/*
查询多行
*/
func queryMultiRowDemo(db *sql.DB) {
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rows, err := db.Query("select ACCOUNT_ID,DISPLAY_NAME from c_account limit 2")
	if err != nil {
		fmt.Printf("未找到结果", err)
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	fmt.Println("查询结果:")
	for rows.Next() {
		var u user
		err := rows.Scan(&u.accountId, &u.displayName)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Println("多行查询结果:"+"id:%d name:%s", u.accountId, u.displayName)
	}
}

/*
插入数据
*/
func insertRowDemo(db *sql.DB) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

/*
更新数据
*/
func updateRowDemo(db *sql.DB) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

/*
删除数据
*/
func deleteRowDemo(db *sql.DB) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

/*
定义存储结构体
*/
type user struct {
	accountId   int
	displayName string
}

func main() {
	bb := initBb()
	//单行查询
	queryRowDemo(bb)
	//多行查询
	queryMultiRowDemo(bb)
	//插入数据
	insertRowDemo(bb)
	//更新数据 返回影响行数
	updateRowDemo(bb)
	//删除数据 返回影响行数
	deleteRowDemo(bb)
	//如果数据库连接打开记得关闭
	defer bb.Close()
}
