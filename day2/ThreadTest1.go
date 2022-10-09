package main

import "fmt"

/*
*
通道
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

	c := make(chan int)
	d := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], d)
	x, y := <-c, <-d // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
