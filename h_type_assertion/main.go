package main

import (
	"fmt"
)

func main() {
	// 类型断言
	// 检查i 是否为nil
	// 检查i存储的值是否未某个类型

	// 使用空接口
	var i interface{} = 10

	//var i2 interface{}
	// 类型断言方式一
	t1 := i.(int)
	fmt.Printf("%T\n", i)
	fmt.Println(t1)

	//t3 := i2.(int)
	//fmt.Println(t3)

	fmt.Println("===============")
	//t2 := i.(string)
	//fmt.Println(t2)

	// 类型断言方式二， 不会出发panic
	t4, ok := i.(int)
	fmt.Printf("%d-%t\n", t4, ok)
	t5, ok := i.(string)
	fmt.Printf("%s-%t\n", t5, ok)

}
