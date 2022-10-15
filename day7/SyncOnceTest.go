package main

import (
	"fmt"
	"sync"
)

/*
SyncOnce 高并发场景下只执行一次 如加载文件或者关闭通道
*/

var loadOnce sync.Once

func main() {
	//入参是方法名称
	loadOnce.Do(test)
}

func test() {
	fmt.Println("测试")
}
