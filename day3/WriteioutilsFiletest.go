package main

import (
	"fmt"
	"io/ioutil"
)

/*
*
文件工具 写入字符串
*/
func main() {
	str := "输入的字符串"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("打开文件失败,err", err)
		return
	}
}
