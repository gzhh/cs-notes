package main

import "fmt"

// 半开方式，两个goroutine通过一个channel通信：交替打印奇偶数字
func printEvenOdd() {
	exit := make(chan struct{})
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Println("A", i)
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			if i%2 == 1 {
				fmt.Println("B", i)
			}
		}
		exit <- struct{}{}
	}()
	<-exit
}

// 半开方式，两个goroutine通过一个channel通信：交替打印（错误示例）
func printChar0() {
	exit := make(chan struct{})
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
			fmt.Println("A", i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println("B", i)
		}
		exit <- struct{}{}
	}()
	<-exit
}

// 封闭方式，两个goroutine通过两个channel通信：交替打印
func printChar1() {
	exit := make(chan struct{})
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- struct{}{}
			fmt.Println("A", i)
			<-ch2
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("B", i)
			ch2 <- struct{}{}
		}
		exit <- struct{}{}
	}()
	<-exit
}

// 缓冲模式，两个goroutine通过两个channel通信：交替打印
// 一个channel带buffer，一个不带
func printChar2() {
	exit := make(chan struct{})
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)

	ch1 <- true
	go func() {
		for i := 0; i < 10; i++ {
			if ok := <-ch1; ok {
				fmt.Println("A", i)
				ch2 <- true
			}
		}
	}()
	go func() {
		defer func() { close(exit) }()
		for i := 0; i < 10; i++ {
			if ok := <-ch2; ok {
				fmt.Println("B", i)
				ch1 <- true
			}
		}
	}()
	<-exit
}

// 封闭模式：使用N个协程交替打印1-100
func printNum() {
	chanNum := 3                                // chan数量，即goroutine数量
	chanQueue := make([]chan struct{}, chanNum) // 创建chan slice
	exitChan := make(chan struct{})             // 退出标识
	var num = 1                                 // 需打印的值

	// 创建chan
	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{})
	}

	// 给最后一个chan写一条数据，为了第一次输出是从第1个chan输出
	go func(i int) {
		chanQueue[i] <- struct{}{}
	}(chanNum - 1)

	// 当前chan如果要打印数据，就必须得上一个chan打印完后才能打印
	for i := 0; i < chanNum; i++ {
		var preChan chan struct{} // 上一个goroutine结束才能输出，控制输出顺序
		var curChan chan struct{} // 当前阻塞输出的goroutine
		if i == 0 {
			preChan = chanQueue[chanNum-1]
		} else {
			preChan = chanQueue[i-1]
		}
		curChan = chanQueue[i]

		go func(i int, preChan, curChan chan struct{}) {
			for {
				if num > 100 {
					// 超过100就退出
					exitChan <- struct{}{}
				}
				// 一直阻塞到上一个输出完，控制顺序
				<-preChan
				fmt.Printf("thread%d: %d \n", i, num)
				num++
				// 当前goroutine已输出
				curChan <- struct{}{}
			}
		}(i, preChan, curChan)
	}

	<-exitChan
	fmt.Println("done")
}

func main() {
	// printEvenOdd()

	// printChar0()
	// printChar1()
	// printChar2()

	printNum()
}
