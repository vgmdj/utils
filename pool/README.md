# 协程池

# 例

```
package main

import (
	"time"

	"github.com/vgmdj/utils/pool"

	"github.com/vgmdj/utils/logger"
)

//主函数
func main() {
	amount := 400000
	standardCal(amount)
	poolCal(amount)

}

func standardCal(amount int) {
	start := time.Now()
	sum := 0
	for i := 0; i < amount; i++ {
		for j := i * 10000; j < (i+1)*10000; j++ {
			sum += j
		}
	}

	end := time.Now()
	logger.Info(sum, end.Sub(start))

}

func poolCal(amount int) {
	start := time.Now()

	//创建一个协程池,最大开启3个协程worker
	p := pool.NewPool(8, 1000)

	result := make(chan int, 1000)
	go func() {
		sum := 0
		count := 0
		for v := range result {
			sum += v

			count++
			if count == amount {
				break
			}
		}

		end := time.Now()

		logger.Info(sum, end.Sub(start))
		p.Close()

	}()

	go func() {

		for i := 0; i < amount; i++ {
			t := pool.NewTask(func(args ...interface{}) {
				calculate(args[0].(int), args[1].(int), args[2].(chan int))
			}, i*10000, (i+1)*10000, result)

			p.Push(t)

		}

	}()

	p.Run()
}

func calculate(start, end int, results chan int) {
	var result = 0
	for i := start; i < end; i++ {
		result += i
	}

	results <- result
}




```