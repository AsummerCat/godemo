package main

import "fmt"

type person struct {
	Name string
	age  int
}

/*
*
对象的方法
*/
func (p person) getName() {
	fmt.Println(p.Name)
}
func main() {
	p1 := person{Name: "小明", age: 11}
	p1.getName()
}