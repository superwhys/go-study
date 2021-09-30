package main

import (
	"fmt"
	"time"
)

type Pool struct {
	work chan func() // 任务
	sem chan struct{} // 数量
}

// New 新建一个协程池
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()), // 任务通道
		sem: make(chan struct{}, size), // 控制协程池大小
	}
}

func (p *Pool) worker(task func()) {
	// 若协程退出，则释放一个空间
	defer func() {
		<- p.sem
		fmt.Println("协程结束")
	}()
	// 无限循环用于一直接受任务
	for {
		task()
		task = <- p.work
	}
}

func (p *Pool) NewTask(task func()) {
	// 第一次进入，由于work是无缓冲通道，则会直接跳过，进入第二个case
	// 此时p.sem通道接受到一个标记，可用协程数量减少了一个
	// 等到p.sem通道满了后，则会进入阻塞，此时若所有协程的任务还没有结束，则两个case都陷入阻塞状态
	// 等到有任务完成后，第一个case会继续接受任务
	select {
		case p.work <- task:
		case p.sem <- struct{}{}:
			go p.worker(task)
	}
}

func test() {
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now())
}


func main() {
	pool := New(2)
	for i := 0; i < 4; i++ {
		pool.NewTask(test)
	}

	time.Sleep(time.Second * 5)

}
