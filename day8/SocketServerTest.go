package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
socket服务端
*/

func main() {
	//开启端口服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("监听端口失败", err)
	}

	for {
		//等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败", err)
			continue
		} else {
			fmt.Println("连接成功,连接方式TCP UDP:", conn.RemoteAddr().Network())
			fmt.Println("连接成功,ip地址:", conn.RemoteAddr().String())
		}
		//处理连接的通信
		go process(conn)
	}
}

/*
处理函数
*/
func process(conn net.Conn) {
	//关闭连接
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("读取客户端数据失败", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}
