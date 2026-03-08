package main

import (
	"fmt"
	"time"
)

/*
 * Worker
 */
type worker struct {
	id  int
	err error
}

func (wk *worker) work(workerChan chan<- *worker) (err error) {
	// 任何 Goroutine 只要异常退出或者正常退出都会调用 defer 函数, 所以在 defer 中向 WorkerManager 的 WorkChan 发送通道
	defer func() {
		// 捕获异常信息, 防止 panic 直接退出
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("Panic happened with [%v]", r)
			}
		} else {
			wk.err = err
		}
		// 通知主 Goroutine, 当前子 Goroutine 已经死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start Worker...ID = ", wk.id)

	// 每个 Worker 睡眠一定时间后, panic 退出或者有 Goexit() 退出
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
	}

	panic("worker panic...")

	//runtime.Goexit()

	return err
}

/*
 * WorkerManager
 */

type WorkerManager struct {
	// 用来监控 Worker 是否已经死亡的缓冲 channel
	workerChan chan *worker
	// 一共要监控的 Worker 的数量
	nWorkers int
}

// 创建一个 WorkerManager 对象
func NewWorkerManager(nworkers int) *WorkerManager {
	return &WorkerManager{
		nWorkers:   nworkers,
		workerChan: make(chan *worker, nworkers),
	}
}

// 用协 worker 池，并为每个 Worker 分配一个 ID, 让每个 Worker 进行工作
func (wm *WorkerManager) StartWorkerPool() {
	// 开启一定数量的 Worker
	for i := 0; i < wm.nWorkers; i++ {
		wk := &worker{id: i}
		go wk.work(wm.workerChan)
	}

	// 启动保活监控
	wm.KeepLiveWorkers()
}

// 保活监控 Workers
func (wm *WorkerManager) KeepLiveWorkers() {
	// 如果有 Worker 已经死亡，则 workerChan 会得到具体死亡的 Worker，然后输出异常，最后重启
	for wk := range wm.workerChan {
		// log the error
		fmt.Printf("Worker %d stopped with err: [%v]\n", wk.id, wk.err)
		// reset err
		wk.err = nil
		// 当检测到一个 wk 已经死亡了，需要重新启动它的业务
		go wk.work(wm.workerChan)
	}
}

func main() {
	wm := NewWorkerManager(10)
	wm.StartWorkerPool()
}
