package main

import "fmt"

/*
func 函数名(形参)(返回值){
	函数体
}
 */

func sum(a int, b int) int{
	return a + b
}

// sums 可变参数传参---多个类型一致的参数
func sums(args ...int) int{
	var sum int
	fmt.Println("参数是：", args)
	for _, v := range args{
		sum += v
	}
	return sum
}

// MyPrint 可变参数传参---多个类型不一致的参数
func MyPrint(args ...interface{}) {
	fmt.Println("参数是：", args)
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an Int value")
		case string:
			fmt.Println(arg, "is an string value")
		case int64:
			fmt.Println(arg, "is an Int64 value")
		default:
			fmt.Println(arg, "is an unknown value")
		}
	}
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}


func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sums(1, 2, 3, 4, 5))
	MyPrint(1, 3.14, "hello world")

	a := []int{1, 2, 3, 4}
	fmt.Println(sums(a...))

	// 匿名函数
	visit(a, func(v int) {
		fmt.Println(v)
	})
}
