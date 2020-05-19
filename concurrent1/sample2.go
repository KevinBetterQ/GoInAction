package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg2     sync.WaitGroup
	mutex   sync.Mutex
)

// 通过互斥锁来控制共享资源的访问
func RunSample2() {
	wg2.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			value++
			counter = value

		}
		mutex.Unlock()
	}
}
