package main

import "fmt"

/*
function types 表示着所有拥有同样的入参类型和返回值类型的函数集合
作为function types的对象， 它拥有着function types的所有方法
 */

type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	greet := Greeting(english)
	greet.say("why")
}
