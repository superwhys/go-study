package main

import (
	"fmt"
	"time"
)

/*
func() // 执行函数
go func() // 开启一个协程执行函数
 */

func myTest() {
	fmt.Println("hello world")
}

func myGo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	//go myTest()
	//fmt.Println("this is main")
	//time.Sleep(time.Second)

	go myGo("routine A")
	go myGo("routine B")
	time.Sleep(time.Second)
}
