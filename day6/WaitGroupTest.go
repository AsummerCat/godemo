package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go test()
	go test()
	test()
	wg.Wait() //等待计数器减为0
	fmt.Println("结束")
}

func test() {
	//新增计数器1
	wg.Add(1)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("输出完成")
	//减少计数器
	defer wg.Done()
}
