// runner package manage the whole life of task
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在规定的时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// 通道,报告从操作系统发送的信号
	interrupt chan os.Signal

	// 通道，报告处理任务已经完成
	complete chan error

	// 通道，报告处理任务已经超时
	timeout <-chan time.Time

	// 一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout receive when task is timeout
var ErrTimeout = errors.New("received timeout")

// ErrorInterrupt receive when task is interrupt
var ErrInterrupt = errors.New("received interrupt")

// New return a new Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add will add task to Runner's tasks
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// run execute all task
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// if detect interrupt signal
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// execute task
		task(id)
	}
	return nil
}

// gotInterrupt check if receiving interrupt signal
func (r *Runner) gotInterrupt() bool {
	select {
	// when interrupt event appear
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// continue run
	default:
		return false
	}
}

// Start execute all tasks and monitor the event of channel
func (r *Runner) Start() error {
	// start receiving all interrupt signal
	signal.Notify(r.interrupt, os.Interrupt)

	// execute tasks
	go func() {
		r.complete <- r.run()
	}()

	select {
	// will execute when tasks completed
	case err := <-r.complete:
		return err

	// execute when timeout
	case <-r.timeout:
		return ErrTimeout
	}
}
