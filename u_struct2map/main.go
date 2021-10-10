package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func struct2map() {
	u1 := User{
		Name: "Superyong",
		Age:  18,
	}

	b, _ := json.Marshal(u1)
	fmt.Printf("%s\n", b)

	var m map[string]interface{}

	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", m)
	for k, v := range m{
		fmt.Printf("key:%v value:%v value_type:%T\n", k, v, v)
	}
}

// 上面的写法int类型的会被转成float64
// toMap 利用反射可以避免上面出现的问题
func toMap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {  // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func main() {
	struct2map()

	u1 := User{
		Name: "SuperYong",
		Age:  18,
	}
	out, _ := toMap(&u1, "json")
	fmt.Println(out)
}
