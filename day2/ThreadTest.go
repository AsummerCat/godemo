package main

import (
	"fmt"
	"time"
)

/*
*
携程
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	go say("xxx")
	say("hello")
}
