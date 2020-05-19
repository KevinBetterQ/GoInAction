package main

import "fmt"

func main() {
	// 主要是 sync.WaitGroup 的使用来控制 goroutine 的结束
	fmt.Println("\n\n…… Run Sample1 ……")
	RunSample1()

	// 加入 sync.Mutex 来控制共享资源的互斥
	fmt.Println("\n\n…… Run Sample2 ……")
	RunSample2()

	// 加入 channel 来管理同步一组 goroutine
	fmt.Println("\n\n…… Run Sample3 ……")
	RunSample3()
}
