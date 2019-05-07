//package pool goroutine pool
//example
//
//package main
//
//import (
//	"time"
//
//	"github.com/vgmdj/utils/pool"
//
//	"github.com/vgmdj/utils/logger"
//)
//
////主函数
//func main() {
//	amount := 400000
//	standardCal(amount)
//	poolCal(amount)
//
//}
//
//func standardCal(amount int) {
//	start := time.Now()
//	sum := 0
//	for i := 0; i < amount; i++ {
//		for j := i * 10000; j < (i+1)*10000; j++ {
//			sum += j
//		}
//	}
//
//	end := time.Now()
//	logger.Info(sum, end.Sub(start))
//
//}
//
//func poolCal(amount int) {
//	start := time.Now()
//
//	//创建一个协程池,最大开启3个协程worker
//	p := pool.NewPool(8, 1000)
//
//	result := make(chan int, 1000)
//	go func() {
//		sum := 0
//		count := 0
//		for v := range result {
//			sum += v
//
//			count++
//			if count == amount {
//				break
//			}
//		}
//
//		end := time.Now()
//
//		logger.Info(sum, end.Sub(start))
//		p.Close()
//
//	}()
//
//	go func() {
//
//		for i := 0; i < amount; i++ {
//			t := pool.NewTask(func(args ...interface{}) {
//				calculate(args[0].(int), args[1].(int), args[2].(chan int))
//			}, i*10000, (i+1)*10000, result)
//
//			p.Push(t)
//
//		}
//
//	}()
//
//	p.Run()
//}
//
//func calculate(start, end int, results chan int) {
//	var result = 0
//	for i := start; i < end; i++ {
//		result += i
//	}
//
//	results <- result
//}


package pool

// Task task
type Task struct {
	args []interface{}
	f    func(args ...interface{})
}

// NewTask init and return a task
func NewTask(f func(args ...interface{}), args ...interface{}) *Task {
	t := Task{
		args: args,
		f:    f,
	}

	return &t
}

// Pool pool struct
type Pool struct {
	// cacheLimit the limit of cacheQueue
	cacheLimit int

	// workers the num of goroutine
	workers int

	// cacheQueue task cache queue
	cacheQueue chan *Task

	// jobsQueue task jobs queue
	jobsQueue chan *Task
}

// NewPool return a pool
func NewPool(workers, cacheLimit int) *Pool {
	return &Pool{
		cacheLimit: cacheLimit,
		workers:    workers,
		cacheQueue: make(chan *Task, cacheLimit),
		jobsQueue:  make(chan *Task, cacheLimit),
	}

}

// worker do the job
func (p *Pool) worker() {
	//worker不断的从jobsQueue内部任务队列中拿任务
	for task := range p.jobsQueue {
		//如果拿到任务,则执行task任务
		task.f(task.args[:]...)
	}
}

// Push receive the task
func (p *Pool) Push(t *Task) {
	p.cacheQueue <- t
}

// Close close the pool
func (p *Pool) Close() {
	close(p.jobsQueue)
	close(p.cacheQueue)
}

// Run start the tasks
func (p *Pool) Run() {
	for i := 0; i < p.workers; i++ {
		go p.worker()
	}

	for task := range p.cacheQueue {
		p.jobsQueue <- task
	}

}
