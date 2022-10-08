package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	strings := []string{"google", "runoob"}
	for index, s := range strings {
		fmt.Println(index, s)
	}

	//无限循环
	for {
		sum++ // 无限循环下去
	}
}
