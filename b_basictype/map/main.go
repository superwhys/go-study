package main

import "fmt"

func main() {
	// 声明并初始化
	// 第一种方法
	var scores1 map[string]int = map[string]int{"english": 80, "chinese": 85}
	fmt.Printf("type is: %T, values is: %v\n", scores1, scores1)
	// 拆分开来就是
	var scores1Split map[string]int
	scores1Split = make(map[string]int)
	scores1Split["english"] = 80
	fmt.Printf("type is: %T, values is: %v\n", scores1Split, scores1Split)

	// 第二种方法
	scores2 := map[string]int{"english": 80, "chinese": 85}
	fmt.Printf("type is: %T, values is: %v\n", scores2, scores2)

	// 第三种方法
	scores3 := make(map[string]int)
	scores3["english"] = 80
	scores3["chinese"] = 85
	fmt.Printf("type is: %T, values is: %v\n", scores3, scores3)

	// 字典中不存在的key默认值是0， 但不能通过值是0去判断key存不存在
	// 判断key存不存在
	value, ok := scores3["math"]
	fmt.Printf("math value is : %v, 存在: %v\n", value, ok)

	// 对字典进行循环
	// 1. 获取key和value
	for key, val := range scores3 {
		fmt.Printf("key: %s, value: %d\n", key, val)
	}

	// 2. 只获取key
	for key := range scores3 {
		fmt.Printf("key: %s\n", key)
	}

	// 3. 只获取value
	for _, val := range scores3 {
		fmt.Printf("value: %d\n", val)
	}
}
