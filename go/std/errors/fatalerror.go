package main

import "fmt"

/*
fatal error 一般是指由运行时（runtime）引发的严重错误，这类错误会导致程序直接崩溃退出。

记录 fatal error 并排查的办法
- GOTRACEBACK来控制Golang panic stack trace输出的信息
- 抓取 coredump 并调试分析
- 参考：https://www.cnblogs.com/apocelipes/p/17536722.html
*/
func main() {
	// fatal error: xxx - deadlock 错误不能被 recover 捕获
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	testWriteMap()

	testDeadlock()
}

func testWriteMap() {
	// fatal error: concurrent map read and map write
	c := make(map[string]int)
	go func() {
		for j := 0; j < 1000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {
		for j := 0; j < 1000; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()
}

func testDeadlock() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	ch <- 1
}
