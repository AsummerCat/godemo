package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
写入文件
*/
func main() {
	fileObj, err := os.OpenFile("./main.go", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败,err", err)
		return
	}
	//写入
	writer := bufio.NewWriter(fileObj)
	writer.Write([]byte("输入的内容"))
	writer.WriteString("输入的字符串")
	//将缓存中的内容写入文件
	writer.Flush()
	defer fileObj.Close()

}
