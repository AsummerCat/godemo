package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
*
文件工具类
*/
func main() {
	//使用ioutil打开文件
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("打开文件失败,err", err)
		return
	}
	fmt.Println(string(content))

	//新版本API
	content, err1 := os.ReadFile("./main.go")
	if err1 != nil {
		fmt.Println("打开文件失败,err", err)
		return
	}
	fmt.Println(string(content))

}
