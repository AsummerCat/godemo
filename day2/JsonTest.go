package main

import (
	"encoding/json"
	"fmt"
)

/*
*
注意字段大写 才能在别的包中查看,所以json序列化的时候需要变成大写
*/
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "小明",
		Age:  100,
	}
	//序列化
	marshal, err := json.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败:%s", err)
	}

	//转换为字符串
	str := string(marshal)
	fmt.Println(str)

	//反序列化
	var p2 Person
	json.Unmarshal([]byte(str), &p2) //传指针修改P2的赋值
	fmt.Println(p2)
}
