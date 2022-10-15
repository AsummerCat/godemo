package main

import (
	"fmt"
	"io"
	"os"
)

/*
*
打开文件读取字节
*/
func main() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("打开文件失败,err", err)
	}

	//指定文件读取长度
	var tmp = make([]byte, 128)

	//循环读取
	for {
		//文件读取 返回读取的字节数和可能的具体错误 如文件末尾会返回 0和io.EOF
		read, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}
		if err != nil {
			fmt.Println("文件读取失败:", err)
			return
		}
		fmt.Println("读取了%d个字节数据", read)
		fmt.Println(string(tmp))
		if read < 128 {
			return
		}
	}
	//关闭文件 函数结束自动调用defer
	defer file.Close()
}
