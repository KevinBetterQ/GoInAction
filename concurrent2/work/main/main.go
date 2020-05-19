package main

import (
	"GoInAction/concurrent2/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"tom",
	"kevin",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Print(n.name)
	time.Sleep(time.Second)
}

var maxGroutines = 2
var tasks = 100

func main() {
	var wg sync.WaitGroup
	wg.Add(tasks * len(names))

	p := work.New(maxGroutines)

	for i := 0; i < tasks; i++ {
		for _, name := range names {
			np := namePrinter{name: name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
