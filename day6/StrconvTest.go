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
		fmt.Println("字符串转换为数字类型转换失败:", err)
	}
	fmt.Println(parseInt)

	//数字类型转换为字符串
	i := int32(97)
	sprintf := fmt.Sprintf("%d", i)
	fmt.Println(sprintf)

	//String转换为int
	atoi, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("String转换为int转换失败:", err)
	}
	fmt.Println(atoi)

	//int转换为String
	itoa := strconv.Itoa(1001)
	fmt.Println(itoa)
}
