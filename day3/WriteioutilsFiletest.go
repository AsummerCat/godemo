package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "输入的字符串"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("打开文件失败,err", err)
		return
	}
}
