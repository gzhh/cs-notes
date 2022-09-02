package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	// 没设置时，交替打印 A 和 B，A 一直是 10；多个协程执行顺序随机
	// 设置为1时，多个协程的执行顺序是固定不变的；具体实现看源码 src/runtime/proc.go 调度逻辑
	// 大致逻辑为：每增加一个子协程就把其对应的函数地址存放到 _p_.runnext，而把 _p_.runnext 原来的地址（即上一个子协程对应的函数地址）移动到队列 _p_.runq 里面，这样当主逻辑执行到 wg.Wait() 时，_p_.runnext 存放的就是最后一个子协程对应的函数地址（即输出B: ９的那个子协程）。当开始执行子协程对应的函数时，首先执行 _p_.runnext 对应的函数，然后按先进先出的顺序执行队列 _p_.runq 里的函数。
	// PS：测试出来先打印 B:9 和 B:8 B:9 这两种情况都有，后一种情况具体原因还待分析。

	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
