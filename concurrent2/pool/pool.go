package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manage many resources which implement io.Close interface
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("pool closed")

// New function create a new Pool
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire get resource from pool
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// check if having resource
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release function release resource to pool
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	// try to put the resource into queue
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// if queue is full, close the resource
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close make pool stop working and close all resources
func (p *Pool) Close() {
	// make operation thread safe
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// first close the channel, then close resources
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}

}
