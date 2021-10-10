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
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	data, _ := proto.Marshal(&p1)

	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)

	err := ioutil.WriteFile("./proto.dat", data, 0644)
	if err != nil {
		return
	}

	//data2, err := ioutil.ReadFile("./proto.dat")
	//if err != nil {
	//	fmt.Printf("read file failed, err:%v\n", err)
	//	return
	//}

	var p2 person.Person

	err = proto.Unmarshal(data, &p2)
	if err != nil {
		return
	}
	fmt.Println(p2)
}
