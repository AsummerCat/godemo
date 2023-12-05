package main

import (
	"fmt"
	"net"
	"time"
)

/*
*
socket客户端
*/
func main() {
	//1.与server建立连接
	client, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接127.0.0.1:20000失败", err)
	} else {
		fmt.Println("连接127.0.0.1:20000成功,连接方式:", client.RemoteAddr().Network())
		fmt.Println("连接127.0.0.1:20000成功,ip地址:", client.RemoteAddr().String())
	}

	for {
		//2.发送数据
		client.Write([]byte("hello world"))

		time.Sleep(time.Duration(2) * time.Second)
	}
	//3.关闭连接
	client.Close()
}
