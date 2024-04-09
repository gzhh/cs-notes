## Mutex
- https://go.dev/tour/concurrency/9
- https://pkg.go.dev/sync#Mutex
- https://go.dev/src/sync/mutex.go

### 临界区
在并发编程中，如果程序中的一部分会被并发访问或修改，那么为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，这部分被保护起来的程序，就叫做临界区。

### 互斥锁 Mutex

我们可以使用互斥锁，限定临界区只能同时由一个线程持有。

互斥锁就很好地解决了资源竞争问题，有人也把互斥锁叫做排它锁。

```
// Locker 的接口定义了锁同步原语的方法集
type Locker interface {
    Lock()
    Unlock()
}

// Mutex 实现了 Locker 接口
type Mutex struct {
	state int32
	sema  uint32
}

func (m *Mutex) Lock()
func (m *Mutex) Unlock()
```

Go 用 Mutex 实现了互斥锁。

互斥锁 Mutex 就提供两个方法 Lock 和 Unlock：进入临界区之前调用 Lock 方法，退出临界区的时候调用 Unlock 方法：

当一个 goroutine 通过调用 Lock 方法获得了这个锁的拥有权后， 其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，直到锁被释放并且自己获取到了这个锁的拥有权。


### Go race detector
- https://go.dev/blog/race-detector

原理：
- 编译器通过探测所有的内存访问，加入代码能监视对这些内存地址的访问（读还是写）。在代码运行的时候，race detector 就能监控到对共享变量的非同步访问，出现 race 的时候，就会打印出警告信息。
- 这个警告不但会告诉你有并发问题，而且还会告诉你哪个 goroutine 在哪一行对哪个变量有写操作，同时，哪个 goroutine 在哪一行对哪个变量有读操作，就是这些并发的读写访问，引起了 data race。

实现机制
- 通过在编译的时候插入一些指令，在运行时通过这些插入的指令检测并发读写从而发现 data race 问题，就是这个工具的实现机制。


### Mutex 使用技巧
- 基本用法
    ```
    package main

    import (
        "fmt"
        "sync"
    )

    func main() {
        var mu sync.Mutex
        var count = 0

        var wg sync.WaitGroup
        wg.Add(10)

        for i := 0; i < 10; i++ {
            go func() {
                defer wg.Done()
                for j := 0; j < 100000; j++ {
                    mu.Lock()
                    count++
                    mu.Unlock()
                }
            }()
        }
        wg.Wait()
        fmt.Println(count)
    }

    ```
- 变量声明
    ```
    // Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可。
    var mu sync.Mutex
    ```
- 嵌套到其他 struct 中
    ```
    type Counter struct {
        mu    sync.Mutex
        Count uint64
    }
    ```
- 如果嵌入的 struct 有多个字段，我们一般会把 Mutex 放在要控制的字段上面，然后使用空格把字段分隔开来。
    ```
    type Counter struct {
        CounterType int
        Name        string

        mu    sync.Mutex
        count uint64
    }
    ```
- 还可以把获取锁、释放锁、计数加一的逻辑封装成一个方法，对外不需要暴露锁等逻辑。
    ```
    // 线程安全的计数器类型
    type Counter struct {
        CounterType int
        Name        string

        mu    sync.Mutex
        count uint64
    }

    // 加1的方法，内部使用互斥锁保护
    func (c *Counter) Incr() {
        c.mu.Lock()
        c.count++
        c.mu.Unlock()
    }

    // 得到计数器的值，也需要锁保护
    func (c *Counter) Count() uint64 {
        c.mu.Lock()
        defer c.mu.Unlock()
        return c.count
    }
    ```

### Mutex 实现原理

- “初版”的 Mutex 使用一个 flag 来表示锁是否被持有，实现比较简单；
- 后来照顾到新来的 goroutine，所以会让新的 goroutine 也尽可能地先获取到锁，这是第二个阶段，我把它叫作“给新人机会”；
- 那么，接下来就是第三阶段“多给些机会”，照顾新来的和被唤醒的 goroutine；
- 但是这样会带来饥饿问题，所以目前又加入了饥饿的解决方案，也就是第四阶段“解决饥饿”。

![](https://static001.geekbang.org/resource/image/c2/35/c28531b47ff7f220d5bc3c9650180835.jpg)

1.初版实现
```
   // CAS操作，当时还没有抽象出atomic包
    func cas(val *int32, old, new int32) bool
    func semacquire(*int32)
    func semrelease(*int32)
    // 互斥锁的结构，包含两个字段
    type Mutex struct {
        key  int32 // 锁是否被持有的标识
        sema int32 // 信号量专用，用以阻塞/唤醒goroutine
    }

    // 保证成功在val上增加delta的值
    func xadd(val *int32, delta int32) (new int32) {
        for {
            v := *val
            if cas(val, v, v+delta) {
                return v + delta
            }
        }
        panic("unreached")
    }

    // 请求锁
    func (m *Mutex) Lock() {
        if xadd(&m.key, 1) == 1 { //标识加1，如果等于1，成功获取到锁
            return
        }
        semacquire(&m.sema) // 否则阻塞等待
    }

    func (m *Mutex) Unlock() {
        if xadd(&m.key, -1) == 0 { // 将标识减去1，如果等于0，则没有其它等待者
            return
        }
        semrelease(&m.sema) // 唤醒其它阻塞的goroutine
    }
```
- CAS 是实现互斥锁和同步原语的基础
    - CAS 指令将给定的值和一个内存地址中的值进行比较，如果它们是同一个值，就使用新值替换内存地址中的值，这个操作是原子性的。
    - 原子性保证这个指令总是基于最新的值进行计算，如果同时有其它线程已经修改了这个值，那么，CAS 会返回失败。
- Mutex 结构体包含两个字段：
    - 字段 key：是一个 flag，用来标识这个排外锁是否被某个 goroutine 所持有，如果 key 大于等于 1，说明这个排外锁已经被持有；
    - 字段 sema：是个信号量变量，用来控制等待 goroutine 的阻塞休眠和唤醒。
    - ![](https://static001.geekbang.org/resource/image/82/25/825e23e1af96e78f3773e0b45de38e25.jpg)
- Unlock 方法细节漏洞
    - Unlock 方法可以被任意的 goroutine 调用释放锁，即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。这是因为，Mutex 本身并没有包含持有这把锁的 goroutine 的信息，所以，Unlock 也不会对此进行检查。Mutex 的这个设计一直保持至今。
    - 所以在使用 Mutex 的时候，必须要保证 goroutine 尽可能不去释放自己未持有的锁，一定要遵循“谁申请，谁释放”的原则。
    - 最佳实践：Lock/Unlock 总是成对紧凑出现，不会遗漏或者多调用，代码更少。
        ```
        f.mu.Lock()
        defer f.mu.Unlock()
        ```
- 初版 Mutext 实现的问题
    - 请求锁的 goroutine 会排队等待获取互斥锁。虽然这貌似很公平，但是从性能上来看，却不是最优的。因为如果我们能够把锁交给正在占用 CPU 时间片的 goroutine 的话，那就不需要做上下文的切换，在高并发的情况下，可能会有更好的性能。

