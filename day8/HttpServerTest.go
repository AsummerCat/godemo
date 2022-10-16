package main

import (
	"fmt"
	"net/http"
)

/**
内置http请求工具 编写HTTP服务端 类似springMvc
*/

func main() {
	//创建一个接口
	http.HandleFunc("/test", f1)
	//开启服务器
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("服务器启动异常", err)
	}

}

func f1(w http.ResponseWriter, r *http.Request) {
	str := "<h1>hello go</h1>"
	w.Write([]byte(str))
}
