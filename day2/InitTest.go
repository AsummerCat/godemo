package main

import "fmt"

/*
*
在执行导入包语句会自动调用内部init()方法函数的调用 ,不能在代码中主动调用它
*/
func init() {
	fmt.Println("初始化函数")
}
func main() {
	fmt.Println("start")
}
