package main

import "fmt"

func main() {
	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//不确定长度
	//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//  将索引为 1 和 3 的元素初始化
	//balance := [5]float32{1:2.0,3:7.0}
	fmt.Println(balance)

	//访问数组元素
	var salary int = balance[3]
	fmt.Println(salary)

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
