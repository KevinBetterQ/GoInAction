package main

import (
	"GoInAction/concurrent2/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Start working.")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	err := r.Start()
	if err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
