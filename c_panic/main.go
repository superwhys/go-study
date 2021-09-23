package main

import "fmt"

func funB() {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("this is func_a")
	panic("func_a error")
}

func funA() {
	defer fmt.Println("close")

	//defer func() {
	//	// recover() 可以将捕获到的panic信息打印
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//
	//fmt.Println("this is func_a")
	//panic("func_a error")
	// 若想捕捉异常后可以跳过，需要包装为一个单独的函数，
	funB()
	fmt.Println("this is func_a2")
}


func main() {
	fmt.Println("this is main")
	funA()
	fmt.Println("this is main2")

}
