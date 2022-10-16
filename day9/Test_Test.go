package main

import (
	"fmt"
	"testing"
)

/**
测试函数
*/

func TestName(t *testing.T) {
	//...
	fmt.Println("1")
	t.Errorf("用例失败")
}
