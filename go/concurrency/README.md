# Go Concurrency

## Concurrency Model
- https://en.wikipedia.org/wiki/Concurrency_(computer_science)

CSP
- Communicating sequential processes https://en.wikipedia.org/wiki/Communicating_sequential_processes

Actor model
- https://en.wikipedia.org/wiki/Actor_model


## Best Practice
- Go 并发编程实战课 - 鸟窝 https://www.pursuitking.com/go_2/
- Go 语言设计与实现 - 6.2 同步原语与锁 https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/



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


## Third Party Package
- conc https://github.com/sourcegraph/conc
