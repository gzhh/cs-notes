package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d 正在处理任务:%d\n", id, j)
		results <- j * 2 // 返回结果
		time.Sleep(time.Second)
	}
}

// Worker Pool 用于限制并发数并复用 Goroutine。
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 1. 启动 3 个固定数量的 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 2. 发送 10 个任务
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // 发送完毕，通知 worker 退出循环

	// 3. 收集结果
	for a := 1; a <= 10; a++ {
		<-results
	}
}
