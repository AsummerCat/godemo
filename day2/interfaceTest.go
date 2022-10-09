package main

import "fmt"

// 定义接口
type Phone interface {
	call()
}

type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	//使用接口定义
	var phone Phone
	//接口实现类
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
