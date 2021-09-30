package main

import (
	"context"
	"fmt"
	"time"
)

func test(ctx context.Context, i int) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("monitor: %v, over\n", i)
			return
		default:
			for {
				time.Sleep(time.Second * 2)
				fmt.Printf("监控器%v，正在监控! \n", i)
			}
		}
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go test(ctx, i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("等待结束")
	cancel()

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep( 5 * time.Second)

	fmt.Println("主程序退出！！")
}
