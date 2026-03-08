### Pool
原理
- https://pkg.go.dev/sync#Pool
- https://www.pursuitking.com/go_2/1-10.html

channel 和 sync.WaitGroup
```
import (
	"fmt"
	"sync"
	"time"
)

func demo1() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		ch <- struct{}{}

		go func(i int) {
			defer func() { <-ch }()

			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
}

func demo2() {
	var wg sync.WaitGroup

	ch := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-ch }()

			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()
```

sync.Pool
- 由浅入深聊聊Golang的sync.Pool https://juejin.cn/post/6844903903046320136
- Go sync.Pool 保姆级教程 https://juejin.cn/post/6989798306440282148
- Go sync.Pool and the Mechanics Behind It https://victoriametrics.com/blog/go-sync-pool/
- Let's dive: a tour of sync.Pool internals https://unskilled.blog/posts/lets-dive-a-tour-of-sync.pool-internals/

第三方库 ants
- https://github.com/panjf2000/ants
- 慢聊Golang协程池Ants实现原理 https://mp.weixin.qq.com/s/fZpPkG-C0wZ5Z45H2aUxAA
    ```
    import (
        "fmt"
        "sync"
        "time"

        "github.com/panjf2000/ants/v2"
    )

    func demo1() {
        // 1. 记得在程序结束时关闭池子，释放资源
        defer ants.Release()

        var wg sync.WaitGroup

        // 2. 创建一个容量为 10 的协程池
        // ants.NewPool(并发数)
        pool, _ := ants.NewPool(10)
        defer pool.Release()

        // 任务函数
        task := func(i int) {
            defer wg.Done()
            fmt.Printf("任务 %d 正在运行，当前活跃协程数: %d\n", i, ants.Running())
            time.Sleep(time.Millisecond * 500) // 模拟耗时操作
        }

        // 3. 提交 100 个任务
        for i := 0; i < 100; i++ {
            wg.Add(1)
            i := i // 闭包变量捕获

            // 提交任务到池中，如果池满则会阻塞或根据配置处理
            _ = pool.Submit(func() {
                task(i)
            })
        }

        wg.Wait()
        fmt.Println("所有任务完成！")
    }

    func demo2() {
        var wg sync.WaitGroup

        // 1. 定义一个特定的处理函数
        // 参数必须是 interface{}，内部可以进行类型断言
        processFunc := func(i interface{}) {
            defer wg.Done()
            num := i.(int)
            fmt.Printf("正在处理数字: %d\n", num)
            time.Sleep(time.Millisecond * 500) // 模拟耗时操作
        }

        // 2. 创建带函数的池
        pf, _ := ants.NewPoolWithFunc(5, processFunc)
        defer pf.Release()

        // 3. 提交任务
        for i := 0; i < 20; i++ {
            wg.Add(1)
            // 直接传入参数即可
            _ = pf.Invoke(i)
        }

        wg.Wait()
        fmt.Printf("最终活跃协程数: %d\n", ants.Running())
    }
  ```


MySQL 连接池
- 标准库 database/sql 配置参数
  - SetMaxOpenConns
  - SetMaxIdleConns
  - SetConnMaxLifetime
  - SetConnMaxIdleTime
- Golang连接池的几种实现案例 https://juejin.cn/post/6844904077386596366
- Configuring sql.DB for Better Performance
  - https://www.alexedwards.net/blog/configuring-sqldb
  - https://colobu.com/2020/05/18/configuring-sql-DB-for-better-performance-2020/
- 详解连接池参数设置（边调边看）https://juejin.cn/post/7111500846575124488

Redis 连接池
