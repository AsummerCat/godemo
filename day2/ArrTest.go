package main

import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < len(n); i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < len(n); j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	/*二维数组*/
	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// 显示结果
	fmt.Println(sites)
}
