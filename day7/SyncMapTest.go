package main

import (
	"fmt"
	"sync"
)

/*
SyncMap  并发安全的Map(开箱即用无需make函数初始化) 内置操作方法
*/

var m2 = sync.Map{}

func main() {

	//存储键值对
	m2.Store("key", "value")

	//查询key的value
	value, ok := m2.Load("key")
	if ok {
		fmt.Println(value)
	}

	//删除key
	m2.Delete("key")

}
