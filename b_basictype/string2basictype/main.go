package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("type is: %T, value is:%v\n", b, b)

	var str2 string = "12345"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 16)
	fmt.Printf("type is: %T, value is:%v\n", n1, n1)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 32)
	fmt.Printf("type is: %T, value is:%v\n", f1, f1)
}
