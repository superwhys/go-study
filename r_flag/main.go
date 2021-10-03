package main

import (
	"flag"
	"fmt"
	"os"
)

/*
命令行参数解析：
  1. 简单情况下可以直接使用 os.Args
  2. 对于复杂的情况可以使用 flag
 */

func method1() {
	fmt.Println(os.Args)
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("arg[%d]=%v\n", index, arg)
		}
	}
}

//var name string
//func init() {
//	flag.StringVar(&name, "name", "superyong", "your name") // 该行建议放在init函数中执行
//}

func method2() {
	var name string
	/*
	第一个参数 ：接收值后，存放在哪个变量里，需为指针
	第二个参数 ：在命令行中使用的参数名，比如 --name jack 里的 name
	第三个参数 ：若命令行中未指定该参数值，那么默认值为 jack
	第四个参数：记录这个参数的用途或意义
	 */
	flag.StringVar(&name, "name", "superyong", "your name") // 该行建议放在init函数中执行
	flag.Parse() // 解析参数
	fmt.Println(name)
}

func main() {
	//method1()
	fmt.Println("=============== flag ===============")
	method2()
}
