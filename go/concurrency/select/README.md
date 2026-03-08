### Select
介绍
- The select statement lets a goroutine wait on multiple communication operations.
  - A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
- The default case in a select is run if no other case is ready.
  - Use a default case to try a send or receive without blocking:
- Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

原理
- https://go.dev/src/runtime/select.go
- select 实现原理: https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/
- 使用反射 reflect.Select 操作 Channel: https://www.pursuitking.com/go_2/1-14.html

关键注意点
- 空 select{}：会永久阻塞当前协程（常用于 main 函数防止程序退出）。
- nil 通道：向 nil 通道发送或接收数据会永远阻塞。在 select 中，这会导致对应的 case 永远不会被选中。
- 优先级：Go 的 select 不支持优先级。如果你需要某个通道有更高优先级，可能需要嵌套 select 或使用额外的逻辑控制。

使用场景
```
// 1.多渠道的选择
func TestDemo1(t *testing.T) {
    select {
    case ret := <-retCh1:
        t.Logf("result %s", ret)
    case ret := <-retCh2:
        t.Logf("result %s", ret)
    default:
        t.Error("No one returned")
    }
}

func TestDemo1() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// 模拟异步数据源 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自通道 1 的消息"
	}()

	// 模拟异步数据源 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自通道 2 的消息"
	}()

	// 模拟异步数据源3
	go func() {
		time.Sleep(10 * time.Second)
		ch3 <- "来自通道 3 的消息"
	}()

	for i := 0; i < 3; i++ {
		// select 会阻塞，直到其中一个 case 就绪
		select {
		case msg1 := <-ch1:
			fmt.Println("接收成功:", msg1)
		case msg2 := <-ch2:
			fmt.Println("接收成功:", msg2)
		case msg3 := <-ch3:
			fmt.Println("接收成功:", msg3)
		case <-time.After(1500 * time.Millisecond):
			// 如果 1.5 秒内没有收到任何消息，执行此分支
			fmt.Println("超时了！")
		}
	}
}


// 2.超时控制
func TestDemo2(t *testing.T) {
    select {
    case ret := <-retCh:
        t.Logf("result %s", ret)
    case <-time.After(time.Second * 1):
        t.Error("time out")
    }
}

// 3.非阻塞发送与退出信号
func worker(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("收到停止信号，正在清理资源...")
			return
		default:
			// 模拟正在执行的任务
			fmt.Println("正在处理任务...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

参考：
- https://go.dev/tour/concurrency/5
- https://go.dev/tour/concurrency/6
- https://gobyexample.com/select
