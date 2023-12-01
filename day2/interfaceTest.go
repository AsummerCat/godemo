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
	//在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()。
	//然后我们在main函数里面定义了一个Phone类型变量，
	//并分别为之赋值为NokiaPhone和IPhone
	//使用接口定义
	var phone Phone
	//接口实现类
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
	//两种实现方式 一种 对象{实例参数} 一种new(对象)
	phone = NokiaPhone{}
	phone.call()

	phone = IPhone{}
	phone.call()

}
