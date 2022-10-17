package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
初始化数据库
*/
func initBb1() (db *sql.DB) {
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

// 预处理查询示例
func prepareQueryDemo(db *sql.DB) {
	sqlStr := "select ACCOUNT_ID,DISPLAY_NAME from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user1
		err := rows.Scan(&u.accountId, &u.displayName)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s", &u.accountId, &u.displayName)
	}
}

// 预处理插入示例
func prepareInsertDemo(db *sql.DB) {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

/*
定义存储结构体
*/
type user1 struct {
	accountId   int
	displayName string
}

func main() {
	/*
	   	 数据库	     占位符语法
	      	MySQL	     ?
	      	PostgreSQL	  $1, $2等
	      	SQLite	     ? 和$1
	       Oracle	     :name
	*/

	bb := initBb()
	//预处理查询示例
	prepareQueryDemo(bb)
	//预处理插入示例
	prepareInsertDemo(bb)
	//如果数据库连接打开记得关闭
	defer bb.Close()
}
