# Garbage collector (GC)

## 介绍
- 一文搞懂七种基本的GC垃圾回收算法 https://mp.weixin.qq.com/s/M8R4QPidlCrr6vix4JUWmg
- 垃圾回收的算法与实现 https://book.douban.com/subject/26821357/

## 原理
实现原理
- A Guide to the Go Garbage Collector https://go.dev/doc/gc-guide
- Getting to Go: The Journey of Go's Garbage Collector https://go.dev/blog/ismmkeynote
- Golang三色标记混合写屏障GC模式全分析 https://www.yuque.com/aceld/golang/zhzanb
- Why is Go's Garbage Collection so criticized? https://www.reddit.com/r/golang/comments/z1o2oe/why_is_gos_garbage_collection_so_criticized/
- An overview of memory management in Go https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8
- Go 语言垃圾收集器的实现原理 | Go 语言设计与实现 https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
- 万字长文图解 Go 内存管理分析：工具、分配和回收原理 https://mp.weixin.qq.com/s/rydO2JK-r8JjG9v_Uy7gXg
- 图解Golang垃圾回收机制！https://mp.weixin.qq.com/s/KXxJpT8D3qFhgiaU6OoYnw
- Golang GC核心要点和度量方法 https://mp.weixin.qq.com/s/0EjrQ-MtQtHdKVCsibYYaA
- Go 垃圾回收器指南 https://mp.weixin.qq.com/s/hbh9drP9pPBcOY7hLNeIAQ
- Go 语言的垃圾回收 https://mp.weixin.qq.com/s/0CXU2jxqiDTOFuk_7AxPCA
- 如何在Go中进行内存优化和垃圾收集器管理(GC)？https://mp.weixin.qq.com/s/eAlQexyK72uXOLlzH4qLAw
- 只改一个参数让Golang GC耗时暴降到1/30！https://mp.weixin.qq.com/s/EEDNuhEr0G4SbTzDhVBjbg
- 图解Golang的GC算法 - https://i6448038.github.io/2019/03/04/golang-garbage-collector/
- Go语言中常见100问题-#99 Not understanding how the GC works https://mp.weixin.qq.com/s/fQensm55bnzPKagMEVs8Ag
- 内存泄漏的7种场景 https://mp.weixin.qq.com/s/n998ILpTkCRpQ73gz7UYaQ
- Best practices in Go to improve performance and reduce pressure on Go’s garbage collector https://medium.com/@aditimishra_541/best-practices-in-go-to-improve-performance-and-reduce-pressure-on-gos-garbage-collector-df10904e0244



src/runtime 源码
- Garbage collector (GC). https://go.dev/src/runtime/mgc.go
- Garbage collector: stack objects and stack tracing https://go.dev/src/runtime/mgcstack.go
- Garbage collector: marking and scanning https://go.dev/src/runtime/mgcmark.go
- Garbage collector: sweeping https://go.dev/src/runtime/mgcsweep.go
- Garbage collector: write barriers. https://go.dev/src/runtime/mbarrier.go
- Garbage collector: type and heap bitmaps. https://go.dev/src/runtime/mbitmap_allocheaders.go
- Garbage collector: type and heap bitmaps. https://go.dev/src/runtime/mbitmap_noallocheaders.go
- Scavenging free pages. https://go.dev/src/runtime/mgcscavenge.go
- Garbage collector work pool abstraction. https://go.dev/src/runtime/mgcwork.go
- https://go.dev/src/runtime/mgclimit.go
- https://go.dev/src/runtime/mgcpacer.go


GC 优化方向 - 减小 heap
- Using sync.Pool (for reusing objects instead of initializing every time). Be careful about the implementation, misimplementation causes more harm than benefit.
- Prefer strings.Builder over + concentration
- Try to preallocate slices and maps if their size can be known
- Reduce pointer usage
- Try to avoid large local variables in a function
- Observe compiler decisions via go build -gcflag="-m" ./... output. Take a look at inlining optimisations. (related -l flag)
- Struct data alignment technique, use fieldalignment linter
- Use the bytedance/sonic instead of the encoding/json library for serialization and deserialization

GC 改进进程
- 本质：不断改进 stw 时长和逻辑程序的并发执行的效率
- (v1.3) 1.1 标记清除 + 完全 STW（开始标记前 - 结束清除）
- (v1.3) 1.2 标记清除 + 提前终止 STW （清除开始前结束 STW）
- (v1.5) 2.三色标记法
  - 2.1 白+灰+黑，初始化全白，从根节点广度遍历，遍历中为灰，遍历完为黑，最终回收所有黑色对象）
  - 2.2 STW（开始标记前 - 清除开始前，即只有白色和黑色对象）
- (v1.5) 3.屏障机制
  - 3.1 触发三⾊标记法不安全的必要条件（需同时满足）
    - 条件1：⼀个⽩⾊对象被⿊⾊对象引⽤（⽩⾊被挂在⿊⾊下）。
    - 条件2：灰⾊对象与它之间的可达关系的⽩⾊对象遭到破坏（灰⾊同时丢了该⽩⾊）。
  - 3.1.1 强三⾊不变式
    - 强三⾊不变⾊实际上是强制性地不允许⿊⾊对象引⽤⽩⾊对象，这样就不会出现⽩⾊对象被误删的情况。
  - 3.1.2 弱三⾊不变式
    - 弱三⾊不变式强调，⿊⾊对象可以引⽤⽩⾊对象，但是这个⽩⾊对象必须存在其他灰⾊对象对它的引⽤，或者可达它的链路上游存在灰⾊对象。
  - 3.2 插⼊屏障和删除屏障（写屏障）
    - 为了遵循上述两种⽅式，GC算法演进到两种屏障⽅式
  - 3.2.1 插入屏障
    - a. 插⼊屏障的具体操作是，在A对象引⽤B对象的时候，B对象被标记为灰⾊（将B挂在A下游，B必须被标记为灰⾊）。插⼊屏障实际上是满⾜强三⾊不变式（不存在⿊⾊对象引⽤⽩⾊对象的情况，因为⽩⾊会强制变成灰⾊）。
    - b. ⿊⾊对象的内存槽有两种位置：栈和堆。栈空间的特点是容⼩，但是要求相应速度快，因为函数调⽤弹出会被频繁地使⽤，所以插⼊屏障机制在栈空间的对象操作中不使⽤，⽽仅仅使⽤在堆空间对象的操作中。
    - c. 步骤
      - （不 STW）三色扫描流程，堆上被黑色引用的白色对象会被标记为灰色，而栈空间则不会
      -（STW）额外的扫描，当全部三⾊标记扫描之后，栈上有可能依然存在⽩⾊对象被引⽤的情况，所以要对栈重新进⾏三⾊标记扫描，但这次为了对象不丢失，这次扫描要启动STW暂停，直到栈空间的三⾊标记结束。
      - 执行清除白色对象操作
    - d. 插⼊屏障的⽬的是保证⿊⾊对象插⼊的时候有灰⾊对象对其保护，或者将被插⼊的对象变为灰⾊，插⼊屏障实则是满⾜强三⾊不变式的⼀种表现，这样就不会出现被误删的⽩⾊对象了。
  - 3.2.2 删除屏障
    - a. 删除屏障的具体操作是，被删除的对象，如果⾃身为灰⾊或者⽩⾊，则被标记为灰⾊。删除屏障实际上是满⾜弱三⾊不变式，⽬的是保护灰⾊对象到⽩⾊对象的路径不会断。
    - b. 删除屏障是当⼀个对象的引⽤被摘掉的时候，或者当⼀个对象引⽤被上游替换的时候，该对象被标记为灰⾊。标记为灰⾊的⽬的：被删除的⽩⾊对象，如果⼜被其他的⿊⾊对象引⽤，则被删除对象应被回收掉。
    - c. 步骤
      - （不 STW）三色扫描流程，额外加上逻辑，当⼀个对象的引⽤被摘掉的时候，或者当⼀个对象引⽤被上游替换的时候，该对象被标记为灰⾊。
      - 执行清除白色对象操作
    - d. 删除屏障依旧可以满⾜并⾏状态下的垃圾回收动作，但是这种⽅式的回收精度较低，因为⼀个对象即使被删除了，最后⼀个指向它的指针也依旧可以“活”过这⼀轮，只有等到下⼀轮GC才会被清理掉。
- (v1.8) 4.混合写屏障
  - 4.1 插⼊写屏障和删除写屏障的缺点
    - 插⼊写屏障：结束时需要STW 新扫描栈，标记栈上引⽤的⽩⾊对象的存活。
    - 删除写屏障：回收精度低，GC开始时STW扫描堆栈来记录初始快照，这个过程会保护开始时刻的所有存活对象。
  - 4.2 混合写屏障的优点
    - 避免了对栈Re-scan（ 新扫描）的过程，这也极⼤地减少了STW的时间，同时结合了插⼊写屏障和删除写屏障两者的优点。
  - 4.3 混合写屏障的规则
    - GC开始将栈上的对象全部扫描并标记为⿊⾊（之后不再进⾏第⼆次 复扫描，⽆须 STW）。
    - GC期间，任何在栈上创建的新对象均为⿊⾊。
    - 被删除的对象标记为灰⾊。
    - 被添加的对象标记为灰⾊。
    - PS：
      - 混合写屏障实际上满⾜的是⼀种变形的弱三⾊不变式。
      - 屏障技术不在栈上应⽤，因为要保证栈的运⾏效率。混合写屏障是GC的⼀种屏障机制，所以只是当程序执⾏GC的时候，才会触发这种机制。
  - 4.4 混合写屏障的场景
    - a. 堆删除引⽤，成为栈下游
    - b. 栈删除引⽤，成为栈下游
    - c. 堆删除引⽤，成为堆下游
    - d. 栈删除引⽤，成为堆下游
    - PS：混合写屏障的延迟问题，在⼀定概率情况下，为了去掉STW会有⼀些内存延迟1个周期被回收。等到下一轮GC，对象如果没有外界添加，它们终将会成为⽩⾊垃圾内存⽽被回收。
