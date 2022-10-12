package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("打开文件失败,err", err)
	}

	//使用缓冲流读取文件
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadSlice('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}
		if err != nil {
			fmt.Println("文件读取失败:", err)
			return
		}
		fmt.Println(line)

	}
	//关闭文件 函数结束自动调用defer
	defer file.Close()
}
