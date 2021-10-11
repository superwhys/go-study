package main

import (
	"fmt"
	"gitee.com/superwhys/go-study/v_protobuf/person"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
)

func main() {
	p1 := person.Person{
		Ids:                  123,
		Name:                 "SuerYong",
		Number:               "123456",
	}

	data, _ := proto.Marshal(&p1)

	fmt.Printf("%+v\n", p1)
	fmt.Printf("%v\n", data)

	err := ioutil.WriteFile("./proto.dat", data, 0644)
	if err != nil {
		return
	}

	var p2 person.Person

	err = proto.Unmarshal(data, &p2)
	if err != nil {
		return
	}
	fmt.Println(p2)
}
