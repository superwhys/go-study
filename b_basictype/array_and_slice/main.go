package main

import "fmt"

func main() {
	// ############### 数组 ###############
	// 声明数组
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Printf("type is: %T, values is: %v\n", arr, arr)

	// 声明并初始化
	var arr2 [3]int = [3]int{4, 5, 6}
	fmt.Printf("type is: %T, values is: %v\n", arr2, arr2)
	arr3 := [3]int{7, 8, 9}
	fmt.Printf("type is: %T, values is: %v\n", arr3, arr3)

	// 让数组根据实际情况自己分配空间
	arr4 := [...]int{10, 11, 12}
	arr5 := [...]int{13, 14, 15, 16}
	fmt.Printf("type is: %T, values is: %v\n", arr4, arr4)
	fmt.Printf("type is: %T, values is: %v\n", arr5, arr5)

	// ############### 切片 ###############
	myarr := [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是: %T\n", myarr[0:2], myarr[0:2])
	// 切片构造
	// 1. 对数组进行片段截取
	myarr2 := [5]int{1,2,3,4,5}

	// 【第一种】
	// 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	sli1 := myarr2[1:3]
	fmt.Printf("sli1 的长度为：%d，容量为：%d\n", len(sli1), cap(sli1))
	fmt.Println(sli1)

	// 【第二种】
	// 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	sli2 := myarr2[1:3:5]
	fmt.Printf("sli2 的长度为：%d，容量为：%d\n", len(sli2), cap(sli2))
	fmt.Println(sli2)

	// 2. 声明 -> 初始化 -> 赋值
	var sli3 []int
	sli3 = []int{1, 2, 3}
	fmt.Printf("sli3 的长度为：%d，容量为：%d\n", len(sli3), cap(sli3))
	fmt.Println(sli3)

	// 3. make 构造函数
	sli4 := make([]int, 2, 2)
	sli4[0] = 1
	sli4[1] = 2
	fmt.Printf("sli4 的长度为：%d，容量为：%d\n", len(sli4), cap(sli4))
	fmt.Println(sli4)

	// 4. 给指定位置赋值
	sli5 := []int{2:3}
	fmt.Printf("sli5 的长度为：%d，容量为：%d\n", len(sli5), cap(sli5))
	fmt.Println(sli5)

	// append
	sli6 := []int{1}
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)
	// 插入一个元素
	sli6 = append(sli6, 2)
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)
	// 插入多个元素
	sli6 = append(sli6, 3, 4)
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)
	// 插入一个切片, ...表示解包
	sli6 = append(sli6, []int{7, 8}...)
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)
	// 在第一个位置插入一个元素, 实际上就是往切片[]int{0}后面插入sli6
	sli6 = append([]int{0}, sli6...)
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)
	// 在中间插入一个切片(两个元素)
	fmt.Println(sli6[:5])
	sli6 = append(sli6[:5], append([]int{5,6}, sli6[5:]...)...)
	fmt.Printf("sli6 的长度为：%d，容量为：%d\n", len(sli6), cap(sli6))
	fmt.Println(sli6)

}
