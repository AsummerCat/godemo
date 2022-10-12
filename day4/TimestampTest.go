package main

import (
	"fmt"
	"time"
)

/*
*
时间戳
*/
func main() {
	timestampDemo()
}

func timestampDemo() {
	now := time.Now() //获取当前时间
	fmt.Println(now)

	unix := now.Unix()     //时间戳
	nano := now.UnixNano() //纳秒时间戳

	fmt.Println(unix)
	fmt.Println(nano)

	//将时间戳转换为时间
	nowDate := time.Unix(unix, 0)
	fmt.Println(nowDate)
}
