package main

import "sync"

/**
sync.WaitGroup 是 Golang 中常用的并发措施，我们可以用它来等待一批 Goroutine 结束。

WaitGroup 的用法非常简单：
使用 Add 添加需要等待的个数，使用 Done 来通知 WaitGroup 任务已完成，使用 Wait 来等待所有 goroutine 结束。

Ref: https://mp.weixin.qq.com/s/CkSd2aldYaoLbd-IKhkpWg
*/
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("hello")
		}()
	}

	wg.Wait()
}
