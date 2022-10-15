package main

import (
	"fmt"
	"time"
)

/*
*
时间类的函数
*/
func main() {
	timeDemo()
}

func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Println(now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒

	fmt.Println(year, month, day, hour, minute, second)
}
