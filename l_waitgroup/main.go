package main

import (
	"fmt"
	"sync"
	"time"
)

/*
由于写的 demo 都是比较简单的， sleep 个 1 秒，我们主观上认为是够用的
但，使用time.Sleep 是一种极不推荐的方式

可以通过一下两种方式来标记完成：
	1. 使用信道来标记完成
	2. 使用 WaitGroup
 */

func test1(done chan bool) {
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	done <- true
}

func test2(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("has done")
		wg.Done()
	}()
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		fmt.Printf("this is: %d\n", i)
	}
}

func main() {
	// 1. 使用信道来标记完成
	done := make(chan bool)
	go test1(done)
	<- done

	fmt.Println("##########################")
	// 2. waitGroup
	var done2 sync.WaitGroup

	done2.Add(2)
	go test2(&done2)
	go test2(&done2)

	done2.Wait()
}
