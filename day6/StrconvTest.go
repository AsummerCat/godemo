package main

import (
	"fmt"
	"strconv"
)

/*
使用strconv标准库实现转换
*/
func main() {
	str := "10000"

	//字符串转换为数字类型
	//表示转换为10进制 64位
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
	}
	fmt.Println(parseInt)

	//数字类型转换为字符串
	i := int32(97)
	sprintf := fmt.Sprintf("%d", i)
	fmt.Println(sprintf)
}
