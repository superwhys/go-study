package main

import (
	"fmt"
	"time"
)

/*
	// 声明+make初始化
	var 信道实例 chan 信道类型
	信道实例 = make(chan 信道类型)

	// 或者合并成一句, 定义了信道容量的信道为缓冲信道, 没有定义的则为无缓冲信道
	信道实例 := make(chan 信道类型, 信道容量)

	e.g:
	pipline := make(chan int)

	信道操作：
		1. 发送数据 pipline <- 200
		2. 读取数据 data := <- pipline

	关闭信道
	close(pipline)
	关闭信道后，接收方仍然可以从信道中取到数据，只是接收到的会永远是 0

	// 读取数据的其他返回值， ok表示通道是否被关闭, 关闭为false
	x, ok := <- pipline

	双向信道和单向信道
	默认为双向信道
	单向信道分为 只读和只写信道
	只读: <- chan type
	只写: chan <- type

*/

type Sender = chan <- int
type Reader = <- chan int

func main() {
	var pipline2 chan int
	pipline2 = make(chan int, 5)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline2))
	pipline2 <- 1
	pipline2 <- 2
	fmt.Printf("当前通道有 %d 个数据\n", len(pipline2))

	pipline := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline <- 1
	fmt.Printf("当前通道有 %d 个数据\n", len(pipline))

	var pipline3 = make(chan int)

	go func() {
		// 只读信道
		var reader Reader = pipline3
		for data := range reader {
			fmt.Println("信道读取到数据：", data)
		}
	}()

	go func() {
		// 只写信道
		var sender Sender = pipline3
		for i := 1; i < 10; i++ {
			sender <- i
		}
		close(sender)
	}()

	time.Sleep(time.Second)
}
