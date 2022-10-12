package main

import (
	"fmt"
	"time"
)

/*
*
基于时间的操作
*/
func main() {
	//追加时间
	addDemo()
	//求差值
	subDemo()
	//在xxx时间之前
	BeforeDemo()
	//在xxx时间之后
	AfterDemo()
	//时间格式化
	TimeFormat()
	//定时器
	TimeTaskDemo()
}

/*
时间加时间间隔 add
*/
func addDemo() time.Time {
	now := time.Now() //获取当前时间
	//加3小时
	addtime := now.Add(time.Hour * 3)
	fmt.Println(addtime)
	return addtime
}

/*
求两个时间之间的差值
*/
func subDemo() {
	now := time.Now() //获取当前时间
	addTime := addDemo()
	sub := addTime.Sub(now)
	fmt.Println(sub)
}

/*
判断时间是否在xxx之前
*/
func BeforeDemo() {
	now := time.Now() //获取当前时间
	addTime := addDemo()
	before := now.Before(addTime)
	fmt.Println(before)
}

/*
判断时间是否在xxx之前
*/
func AfterDemo() {
	now := time.Now() //获取当前时间
	addTime := addDemo()
	after := now.After(addTime)
	fmt.Println(after)
}

/*
定时器 time.Tick(时间间隔)
*/
func TimeTaskDemo() {
	tick := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range tick {
		fmt.Println(i) //每秒都会执行的任务
	}
}

/*
时间格式化 Format
*/
func TimeFormat() {
	now := time.Now()
	//格式化的模板为Go的出生日期 2006年1月2号15点04分

	//24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//12小时制
	fmt.Println(now.Format("2006-01-02 3:04:05.000 PM Mon Jan"))

	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006/01/02 :15:04:05"))
	fmt.Println(now.Format("15:04 \"2006/01/02"))
}
