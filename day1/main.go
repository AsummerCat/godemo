package main

import "fmt"

// 变量批量声明
var (
	name string
	age  int
	isOk bool
)

// 主函数执行
// 函数外 只能设置标识符(变量,常量,函数,类型)的声明
func main() {
	fmt.Println("hello Word")
	name = "小明"
	age = 18
	isOk = true

}
