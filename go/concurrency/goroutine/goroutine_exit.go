package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	handle1()
}

// 方法1：使用 context.Context 取消机制（推荐）
// context 是处理 goroutine 生命周期的标准方法，适用于复杂并发系统。
// 场景：复杂并发/取消链
func handle1() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker1(ctx)

	time.Sleep(time.Second * 3)
	cancel() // 通知退出
	time.Sleep(time.Second)
}

func worker1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine context cancelled")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Second)
		}
	}
}

// 方法2：使用 channel 通信退出信号
// 使用 chan struct{} 或 chan bool 作为退出信号是一种常用、优雅的方式。
// 场景：控制单个 goroutine
func handle2() {
	stopChan := make(chan struct{})
	go worker2(stopChan)

	time.Sleep(3 * time.Second)
	close(stopChan)             // 通知退出
	time.Sleep(1 * time.Second) // 等待 goroutine 打印退出日志
}

func worker2(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Goroutine received stop signal")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// 方法3：使用 sync.WaitGroup 管理并等待退出
// 配合 channel / context.WithCancel
// 场景：等待所有 goroutine 退出
func handle3() {
	stopChan := make(chan struct{})
	wg.Add(1)
	go worker3(stopChan, &wg)

	time.Sleep(3 * time.Second)
	close(stopChan)
	wg.Wait()
}

var wg sync.WaitGroup

func worker3(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker done")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
