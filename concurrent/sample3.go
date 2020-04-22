package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 的数量
	taskLoad         = 10 // 要处理的工作的数量
)

var wg3 sync.WaitGroup

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// 管理一组 goroutine 来接收并完成工作
func RunSample3() {
	// 创建一个有缓冲的管道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutines 来处理工作
	wg3.Add(numberGoroutines)
	for gr := 0; gr < numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完成时关闭通道
	// 以便所有 goroutine 退出
	close(tasks)

	wg3.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg3.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示开始工作
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// 显示完成工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
