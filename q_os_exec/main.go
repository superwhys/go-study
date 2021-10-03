package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

/*
执行命令的库是 os/exec，exec.Command 函数返回一个 Cmd 对象，根据不同的需求，可以将命令的执行分为三种情况
1. 只执行命令，不获取结果
2. 执行命令，并获取结果（不区分 stdout 和 stderr）
3. 执行命令，并获取结果（区分 stdout 和 stderr）
*/

// 1. 只执行命令，不获取结果
func method1() {
	cmd := exec.Command("ls", "-l")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// 2. 执行命令，并获取结果（不区分 stdout 和 stderr）
func method2() {
	cmd := exec.Command("ls", "-l")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

// 3. 执行命令，并获取结果（区分 stdout 和 stderr）
func method3() {
	cmd := exec.Command("ls", "-l", "/var/log/*.log")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout  // 标准输出
	cmd.Stderr = &stderr  // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}


func main() {
	method1()
	method2()
	method3()
}