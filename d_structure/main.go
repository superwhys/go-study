package main

import "fmt"

// 定义结构体
type person struct {
	name, gender string
	age int
	mother *person
	father *person
}

// 结构体绑定方法
func (per person) callName() {
	fmt.Println(per.name)
}

func main() {
	// 结构体实例化
	var Ming person
	// 结构体实例化方法2
	MingFather := person{
		name: "XiaoQiang",
		gender: "man",
	}
	// 结构体实例化方法2
	// 等价于 var MingMother *person = new(person)
	MingMother := new(person)
	MingMother.name = "XiaoHong"
	MingMother.gender = "male"

	Ming.name = "XiaoMing"
	Ming.gender = "man"
	Ming.father = &MingFather
	Ming.mother = MingMother
	fmt.Printf("type is: %T, values is: %v\n", Ming, Ming)
	fmt.Printf("type is: %T, values is: %v\n", MingFather, MingFather)
	fmt.Printf("type is: %T, values is: %v\n", MingMother, MingMother)

	fmt.Printf("Ming Father is: %v\n", *Ming.father)

	// 结构体绑定方法
	Ming.callName()
}
