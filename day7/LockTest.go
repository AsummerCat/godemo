package main

import (
	"fmt"
	"sync"
)

// 创建一个全局锁
var lock sync.Mutex

func main() {
	tryLock := lock.TryLock()
	if tryLock {
		fmt.Println("已锁定")
	} else {
		fmt.Println("未锁定")
	}
	lock.Lock()

	lock.Unlock()
}
