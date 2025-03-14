# Go Concurrency

## Concurrency Model
- https://en.wikipedia.org/wiki/Concurrency_(computer_science)

CSP
- Communicating sequential processes https://en.wikipedia.org/wiki/Communicating_sequential_processes

Actor model
- https://en.wikipedia.org/wiki/Actor_model


## Best Practice
- Go 并发编程实战课 - 鸟窝 https://www.pursuitking.com/go_2/
  - 为什么说并发编程很难？https://colobu.com/2023/07/16/why-conucrrent-programming-is-hard/
  - 深入理解Go并发编程 https://book.douban.com/subject/36667173/
- Concurrency in Go https://book.douban.com/subject/26994591/
  - Concurrency in Go 中文笔记 https://www.kancloud.cn/mutouzhang/go/596804
- Go 语言设计与实现 - 6.2 同步原语与锁 https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/
- Golang异步编程方式和技巧 https://mp.weixin.qq.com/s/PuNu65ggHyB6jxRqhbN_VQ
- 一个例子，给你讲透典型的Go并发控制 https://mp.weixin.qq.com/s/qxs6ikRy_Km5JlJesYmH4w
- 20种不同并发模型示例，带你深入理解并发模型 https://mp.weixin.qq.com/s/aTYie9PHcoXRYYmmIGaOXA
- Goroutines in Go: A Practical Guide to Concurrency https://getstream.io/blog/goroutines-go-concurrency-guide/

### Concurrent Servers
- https://eli.thegreenplace.net/2017/concurrent-servers-part-1-introduction/

### 问题研究
- Understanding Real-World Concurrency Bugs in Go https://songlh.github.io/paper/go-study.pdf


## Concurrency Patterns
- Go Concurrency Patterns
  - https://go.dev/talks/2012/concurrency.slide#1
  - https://www.youtube.com/watch?v=f6kdp27TYZs
- Advanced Go Concurrency Patterns
  - https://go.dev/talks/2013/advconc.slide#1
  - https://www.youtube.com/watch?v=QDDwwePbDtw
- https://github.com/lotusirous/go-concurrency-patterns
- https://github.com/luk4z7/go-concurrency-guide
- Go 语言高级编程- 1.6常见的并发模式 https://chai2010.cn/advanced-go-programming-book/ch1-basic/ch1-06-goroutine.html

Go Concurrency Patterns: Pipelines and cancellation
- https://go.dev/blog/pipelines

高效I/O并发处理：双缓冲和Exchanger
- https://colobu.com/2024/01/14/double-buffering-and-Exchanger/


## Synchronization Patterns
### sync
- https://pkg.go.dev/sync
  - mutex
  - rwmutex
  - waitgroup
  - once
  - pool
  - cond
  - sync.Map

### x/sync
- https://pkg.go.dev/golang.org/x/sync
  - errgroup
  - semaphore
  - singleflight
  - syncmap

### Ref
- https://victoriametrics.com/blog/go-sync-mutex/
- https://victoriametrics.com/blog/go-sync-waitgroup/
- https://victoriametrics.com/blog/go-sync-pool/
- https://victoriametrics.com/blog/go-sync-cond/
- https://victoriametrics.com/blog/go-singleflight/


## Third Party Package
- conc https://github.com/sourcegraph/conc
- xsync https://github.com/puzpuzpuz/xsync
