package main

import (
	"sync"
)

// 创建一个读写锁
var (
	rwlock sync.RWMutex
)

/*
读写锁
*/
func main() {
	//读锁
	rwlock.RLock()
	//解除读锁
	rwlock.RUnlock()
	//写锁
	rwlock.Lock()
	//解除写锁
	rwlock.Unlock()
}
