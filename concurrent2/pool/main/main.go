// the simulation implement a pool of database connections
package main

import (
	"GoInAction/concurrent2/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 10
	pooledResources = 2
)

// shared resource
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// used to assign a unique id to every connection
var idCounter int32

// createConnection is a factory function to create new dbConnection
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// create pool
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// use pool
	for query := 0; query < maxGoroutines; query++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()

}

// performQueries used to test connection
func performQueries(q int, p *pool.Pool) {
	// acquire a connection from pool
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// release a connection to pool
	defer p.Release(conn)

	// simulate query
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]", q, conn.(*dbConnection).ID)
}
