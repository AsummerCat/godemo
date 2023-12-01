package main

import "fmt"

/*
*
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据

	// 并把值赋给 v

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
ch := make(chan int)
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//创建通道c,d
	c := make(chan int)
	d := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], d)
	x, y := <-c, <-d // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
