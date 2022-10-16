package main

import (
	"fmt"
	"net/http"
)

/**
内置http请求工具 编写HTTP客户端 连接访问
*/

func main() {

	//get请求

	get, err := http.Get("http://www.baidu.com")
	defer get.Body.Close()
	if err != nil {
		fmt.Println("访问失败", err)
	}
	fmt.Println("访问成功,状态码", get.StatusCode)

	//post请求
	post, err := http.Post("http://www.baidu.com", "application/json", nil)
	defer post.Body.Close()
	if err != nil {
		fmt.Println("访问失败", err)
	}
	fmt.Println("访问成功,状态码", post.StatusCode)

}
