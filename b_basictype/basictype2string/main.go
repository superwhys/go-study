package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 90
	var num2 float64 = 3.14
	var b bool = true
	var mychar byte = 'h'
	var str string

	// 使用fmt.Sprintf转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("type is: %T, str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("type is: %T, str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("type is: %T, str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("type is: %T, str=%q\n", str, str)

	// 使用strconv转换
	var num3 int = 90
	var num4 float64 = 3.14
	var b2 bool = true
	var str2 string

	str2 = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("type is: %T, str=%q\n", str2, str2)

	str2 = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("type is: %T, str=%q\n", str2, str2)

	str2 = strconv.FormatBool(b2)
	fmt.Printf("type is: %T, str=%q\n", str2, str2)
}
