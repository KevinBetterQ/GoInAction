// work package manage a pool of goroutines
package work

import (
	"log"
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	workers chan Worker
	wg      sync.WaitGroup
}

// New function create a new pool
func New(maxGroutines int) *Pool {
	p := Pool{
		workers: make(chan Worker),
	}

	p.wg.Add(maxGroutines)
	for i := 0; i < maxGroutines; i++ {
		go func(id int) {
			for w := range p.workers {
				log.Printf("worker[%d] working ……", id)
				w.Task()
			}
			p.wg.Done()
		}(i)
	}

	return &p
}

// Run function submit task to pool
func (p *Pool) Run(w Worker) {
	p.workers <- w
}

// Shutdown function wait for all goroutines close
func (p *Pool) Shutdown() {
	close(p.workers)
	p.wg.Wait()
}
