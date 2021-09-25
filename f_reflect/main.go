package main

import (
	"fmt"
	"reflect"
)

type Profile struct {
	name string
	age int
	gender string
}

type Person struct {
	name string
	age int
	gender string
}

func (p Person)SayBye() string {
	//fmt.Println("Bye")
	return "Bye"
}

func (p Person)SayHello() string {
	//fmt.Println("Hello")
	return "hello"
}

func (p Person)SaySomething(sentence string) {
	fmt.Printf("say %s", sentence)
}


func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()
	fmt.Println("arr反射后的对象值为：", value)
	value.Set(reflect.Append(value, reflect.ValueOf(3)))
	fmt.Println("arr反射append后值为：", value)
	fmt.Println(value.Len())
}


func main() {
	// 反射三定律
	// 1. 反射可以将接口类型变量 转换为“反射类型对象”；
	// 2. 反射可以将 “反射类型对象”转换为 接口类型变量；
	// 3. 如果要修改 “反射类型对象” 其类型必须是 可写的；

	var a int = 10
	// 将接口变量转换为反射对象，
	// 但这里的a是int不是接口类型，
	// 由于 TypeOf 和 ValueOf 两个函数接收的是 interface{} 空接口类型，
	// 而 Go 语言函数都是值传递，因此Go语言会将我们的类型隐式地转换成接口类型。
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("反射对象Type的类型为：%T\n", t)
	fmt.Printf("反射对象Value的类型为：%T\n", v)

	// 反射对象到接口变量的转换, 之后value对象可以转换为接口变量
	i := v.Interface()
	fmt.Printf("类型为：%T, 值为：%v \n", i, i)
	// 此处的i静态类型为空接口类型，若需要转为最初的原始类型，需要进行类型断言

	// 反射类型对象的可写性， 若不可写且还去修改它，会报错
	fmt.Printf("v 可写性: %t\n", v.CanSet()) // 不可写
	//v.SetInt(20) // 会报错

	// 要让反射对象具备可写性，需要注意两点
	// 1. 创建反射对象时传入变量的指针
	// 2. 使用 Elem()函数返回指针指向的数据
	v1 := reflect.ValueOf(&a)
	fmt.Printf("v1 可写性: %t\n", v1.CanSet()) // 不可写
	v1 = v1.Elem()
	fmt.Printf("v1 可写性: %t\n", v1.CanSet()) // 可写
	v1.SetInt(30)
	fmt.Printf("反射修改值后的a为：%d\n", a)

	// Type() 和 Kind()
	// Kind() 表示更基础，范围更广的分类
	// 当使用TypeOf() 时， 则没有Type()函数
	// 而只有使用ValueOf时， 才有Type()函数
	// 定义一个结构体Profile

	// TypeOf
	m := Profile{}
	t2 := reflect.TypeOf(m)
	fmt.Printf("Type: %v, Kind: %v\n", t2, t2.Kind())

	t3 := reflect.TypeOf(&m)
	fmt.Printf("Type: %v, Kind: %v\n", t3, t3.Kind())
	fmt.Printf("Type: %v, Kind: %v\n", t3.Elem(), t3.Elem().Kind())

	fmt.Println("######### ValueOf #########")
	// ValueOf
	t4 := reflect.ValueOf(m)
	fmt.Printf("Type: %v, Kind: %v\n", t4.Type(), t4.Kind())

	t5 := reflect.ValueOf(&m)
	fmt.Printf("Type: %v, Kind: %v\n", t5.Type(), t5.Kind())
	fmt.Printf("Type: %v, Kind: %v\n", t5.Elem().Type(), t5.Elem().Kind())

	arr := []int{1, 2}
	fmt.Println("arr长度为：", len(arr))
	fmt.Println( "arr值为：", arr)
	appendToSlice(&arr)
	fmt.Println("append后的值为：", arr)

	//  反射对属性的操作
	// NumField() 和 Field()
	p := Person{"写代码的明哥", 27, "male"}
	v3 := reflect.ValueOf(p)
	fmt.Println("字段数为：", v3.NumField())

	fmt.Println("第 1 个字段：", v3.Field(0))
	fmt.Println("第 2 个字段：", v3.Field(1))
	fmt.Println("第 3 个字段：", v3.Field(2))

	// 反射对方法的操作
	// NumMethod() 和 Method()
	fmt.Println("方法数为：", v3.NumMethod())
	fmt.Println("第 1 个方法：", v3.Method(0))
	fmt.Println("第 2 个方法：", v3.Method(1))
	// 要获取方法名要使用TyeOf
	t6 := reflect.TypeOf(p)
	fmt.Println("方法数为：", t6.NumMethod())
	fmt.Println("第 1 个方法：", t6.Method(0))
	fmt.Println("第 2 个方法：", t6.Method(1).Name)

	// 动态调用函数
	// 使用Call方法， 要使用ValueOf
	fmt.Println("############ 动态调用函数 ##############")

	// 使用索引， 无参数
	v4 := reflect.ValueOf(p)
	for i := 0; i < v4.NumMethod() - 1; i++ {
		fmt.Printf("调用第%d个方法: %v, 结果为: %v\n",
			i+1,
			t6.Method(i).Name,
			v4.Method(i).Call(nil))
	}

	// 使用函数名，无参数
	res := v4.MethodByName("SayHello").Call(nil)
	fmt.Printf("使用函数名调用函数，结果为: %v\n", res)

	// 使用函数名且有参数
	sentence := reflect.ValueOf("Hello World")
	input := []reflect.Value{sentence}
	v4.MethodByName("SaySomething").Call(input)
}