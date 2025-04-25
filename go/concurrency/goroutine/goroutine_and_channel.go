package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func handle(ch chan struct{}, i int) {
	defer wg.Done()

	fmt.Println("handle , goroutine count = ", i, runtime.NumGoroutine())
	<-ch
}

// buffer chan + sync
func main1() {
	cnt := math.MaxInt
	// cnt := 10
	ch := make(chan struct{}, 5)
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go handle(ch, i)
	}

	wg.Wait()
}

func foo(ch chan int) {
	for i := range ch {
		fmt.Println("handle , goroutine count = ", i, runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(ch chan int, task int) {
	wg.Add(1)
	ch <- task
}

// unbuffer chan + sync
func main() {
	cnt := math.MaxInt
	ch := make(chan int)
	goCnt := 5

	for i := 0; i < goCnt; i++ {
		go foo(ch)
	}

	for i := 0; i < cnt; i++ {
		sendTask(ch, i)
	}

	wg.Wait()
}
