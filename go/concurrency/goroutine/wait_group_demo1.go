package main

import (
	"fmt"
	"sync"
)

/*
// 输出结果为：over!
// 因为 goroutine 以非阻塞的方式执行，它们会随着程序(主线程)的结束而消亡，所以程序输出字符串 "over!" 就退出了。
func main()  {
	names := []string{
		"Zhao",
		"Qian",
		"Sun",
		"Li",
	}

	for _, name := range names {
		go helloHandler(name)
	}
	fmt.Println("over!")
}

func helloHandler(name string) {
	time.Sleep(time.Second)
	fmt.Println("Hello, ", name)
}
 */


// To wait for multiple goroutines to finish, we can use a wait group.
// Ref:
// https://golang.org/pkg/sync/#example_WaitGroup
// https://gobyexample.com/waitgroups
func main()  {
	names := []string{
		"Zhao",
		"Qian",
		"Sun",
		"Li",
	}

	var wg sync.WaitGroup
	for _, name := range names {
		// Add adds delta, which may be negative, to the WaitGroup counter.
		wg.Add(1)
		go helloHandler(name, &wg)
	}
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()

	fmt.Println("over!")
}

// This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.
func helloHandler(name string, wg *sync.WaitGroup) {
	// Done decrements the WaitGroup counter by one.
	defer wg.Done()

	//time.Sleep(time.Second)
	fmt.Println("Hello, ", name)
}
