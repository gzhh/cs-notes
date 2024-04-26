package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// chan + sync 控制 goroutine 并发数量
func main() {
	workerNum := 2
	taskCh := make(chan *Task, workerNum)

	var wg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		i = i
		wg.Add(1)
		go func(taskCh chan *Task) {
			defer wg.Done()

			for task := range taskCh {
				if task.Name == "" {
					// 这里不能使用 return 返回，因为 taskCh 中可能有未处理的任务，直接返回 return 可能导致程序卡死。
					continue
				}

				fmt.Printf("Worker %d, Task %d\n", i, task.ID)
				// doing something
				time.Sleep(time.Second)
			}
		}(taskCh)
	}

	tasks := []*Task{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
	}
	for _, task := range tasks {
		taskCh <- task

	}
	close(taskCh)

	wg.Wait()
}
