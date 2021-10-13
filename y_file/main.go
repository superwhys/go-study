package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func bufioReadFile() {
	file, _ := os.Open("./file.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("has done!")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func ioutilReadFile() {
	fmt.Println("============================")
	content, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func writeFile() {
	file, err := os.OpenFile("./file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "\nhello world\n"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("\nhello world again") //直接写入字符串数据
}

func bufioWrite() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("\n")
	for i := 0; i < 10; i++ {
		writer.WriteString("hello superyong\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

// ioutilWrite 在写文件时若文件存在会清空其原来的内容
func ioutilWrite() {
	str := "hello superyong"
	err := ioutil.WriteFile("./file.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	bufioReadFile()
	ioutilReadFile()
	writeFile()
	bufioWrite()
	ioutilWrite()
}
