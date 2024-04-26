package main

import (
	"fmt"
	"sync"
	"time"
)

// channel + sync 控制并发数量
// 发送1000个请求，并控制10个并发
// sync.WaitGroup + channel
func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)
	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}
