package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Home struct {
	Owner *Person
	Address string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age: 22,
	}

	h1 := Home{
		Owner: &p1,
		Address: "Address",
	}
	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	data2, err := json.Marshal(h1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data1) // {"name":"Jack","age":22,"addr":""}
	fmt.Printf("%s\n", data2) // {"Owner":{"name":"Jack","age":22,"addr":""},"Address":"Address"}

	// 反射
	// 获取field
	p := reflect.TypeOf(Person{})
	fmt.Printf("%v\n", p)
	name, _ := p.FieldByName("Name")
	fmt.Printf("%v\n", name)
	name1 := reflect.ValueOf(Person{}).Type().Field(1)
	fmt.Printf("%v\n", name1)
	name2 := reflect.ValueOf(&Person{}).Elem().Type().Field(2)
	fmt.Printf("%v\n", name2)

	// 获取tag
	tag := name1.Tag
	fmt.Println(tag)

	// 获取键值对
	labelValue := tag.Get("json")
	fmt.Println(labelValue)
}
