package main

import (
	"fmt"
	"sync"
)

/*
使⽤两个goroutine交替打印序列，⼀个goroutine打印数字， 另外⼀个goroutine打印字⺟， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/
func main() {
	number, letter := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			<-number
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- true
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-letter
			if i >= 26 {
				wg.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			number <- true
		}
	}(&wg)

	number <- true
	wg.Wait()
}

// 解法2
func answer2() {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	quit := make(chan interface{})

	i := 0

	// print num
	go func() {
		for {
			<-ch1

			fmt.Printf("%d%d", i+1, i+2)

			ch2 <- struct{}{}
		}
	}()

	// print alphabet
	go func() {
		for {
			<-ch2

			if i >= len(chars) {
				quit <- struct{}{}
				return
			}
			fmt.Printf("%c%c", chars[i], chars[i+1])
			i += 2

			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	<-quit
}
