package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync包
1. Mutex 互斥锁
2. RwMutex 读写锁

 */

// 不加锁会由于并发操作而造成冲突
func add(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count += 1
	}
	wg.Done()
}


func addMutex(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		// 在进行操作前先加锁， 操作完后在释放锁
		lock.Lock()
		*count += 1
		lock.Unlock()
	}
	wg.Done()
}


func noMutexTest() {
	// 利用waitGroup来标记完成
	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	go add(&count, &wg)
	go add(&count, &wg)
	go add(&count, &wg)

	wg.Wait()
	fmt.Println("value of count is: ", count)
}


func mutexTest() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0
	wg.Add(3)
	go addMutex(&count, &wg, lock)
	go addMutex(&count, &wg, lock)
	go addMutex(&count, &wg, lock)

	wg.Wait()
	fmt.Println("value of count is: ", count)
}


func readCount(flag int, count *int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	for i:=0; i<10; i++ {
		lock.RLock()
		fmt.Printf("第%d个用户读取数据的结果是: %d\n", flag, *count)
		time.Sleep(time.Second * 2)
		lock.RUnlock()
	}
	wg.Done()
}

func writeCount(flag int, count *int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	for i:=0; i<10; i++ {
		lock.Lock()
		*count += 1
		fmt.Printf("======第%d个用户写入数据后count是: %d\n", flag, *count)
		lock.Unlock()
	}
	wg.Done()
}

func RWMutexTest() {
	lock := &sync.RWMutex{}
	var wg sync.WaitGroup
	count := 0

	wg.Add(4)
	go writeCount(3, &count, &wg, lock)
	go readCount(1, &count, &wg, lock)
	go writeCount(1, &count, &wg, lock)
	go readCount(2, &count, &wg, lock)
	go readCount(3, &count, &wg, lock)
	go readCount(4, &count, &wg, lock)
	go writeCount(2, &count, &wg, lock)

	wg.Wait()
}



func main() {
	//// 运行后多次结果都不同
	//noMutexTest()
	//
	//// 利用锁机制之后就不会出现结果不同的问题
	//mutexTest()

	RWMutexTest()

}
