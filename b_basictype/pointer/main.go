package main

import "fmt"

func main() {
	// 普通变量
	var num int = 1
	// 指针变量: 1. 声明 -> 初始化 -> 赋值
	// 声明指针
	var pint *int
	//指针的初始值是nil
	fmt.Printf("type is: %T, value is: %v\n", pint, pint)
	// 给声明后的指针分配内存空间
	pint = new(int)
	// 将num值赋值给指针变量指向地址空间的值
	*pint = num
	// 或者 （初始化+赋值）
	//pint = &num
	fmt.Printf("type is: %T, space is: %p, space_value is: %d\n", pint, pint, *pint)

	// 指针变量: 2. 通过 ':=' 语法糖
	pint2 := &num
	fmt.Printf("num space is: %v\n", &num)
	fmt.Printf("type is: %T, space is: %p, space_value is: %d\n", pint2, pint2, *pint2)

	pstr := new(string)
	*pstr = "hello world"
	fmt.Printf("type is: %T, space is: %p, space_value is: %s\n", pstr, pstr, *pstr)


}
