# Memory management

## Memory model
- https://en.wikipedia.org/wiki/Memory_model_(programming)
  - https://en.wikipedia.org/wiki/Memory_ordering
  - https://en.wikipedia.org/wiki/Memory_barrier
  - https://en.wikipedia.org/wiki/Optimizing_compiler
  - https://en.wikipedia.org/wiki/Race_condition

Ref
- 如何系统的学习 Memory Model? https://www.zhihu.com/question/23572082
- 聊聊内存模型与内存序 https://mp.weixin.qq.com/s/t5_Up2YZEZt1NLbvgYz9FQ
- 并发理论基础：指令重排序问题 https://zhuanlan.zhihu.com/p/298448987
- A Primer on Memory Consistency and Cache Coherence https://pages.cs.wisc.edu/~markhill/papers/primer2020_2nd_edition.pdf

### Go Memory Model
- Memory Models https://research.swtch.com/mm
- The Go Memory Model https://go.dev/ref/mem
- Go’s Memory Model http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf
- Go Memory Model https://blog.laisky.com/p/gmm/?lang=zh
- Go 内存模型：happens-before 原则 https://eddycjy.com/posts/go/memory-model/
- 关于 golang 官网一段代码的疑惑 https://v2ex.com/t/1031075
- 一站式Golang内存管理洗髓经 https://www.yuque.com/aceld/golang/qzyivn#EaWZp


Go Race Detector
- Introducing the Go Race Detector https://go.dev/blog/race-detector
- Data Race Detector https://go.dev/doc/articles/race_detector

## 常见内存节省技巧
- 预先给切片分配容量
- 按类型调整结构体中字段顺序-内存对齐
- 使用空结构体 struct{}

## 内存对齐
Data structure alignment
- https://en.wikipedia.org/wiki/Data_structure_alignment
- https://zh.wikipedia.org/wiki/字节序
- https://go101.org/article/memory-layout.html
- 什么是大端序和小端序，为什么要有字节序？https://zhuanlan.zhihu.com/p/352145413
- 一文轻松理解内存对齐 https://cloud.tencent.com/developer/article/1727794
- Go struct 内存对齐 https://geektutu.com/post/hpg-struct-alignment.html
- 在 Go 中恰到好处的内存对齐 https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/go-memory-align
- 【Golang】这个内存对齐呀 https://www.bilibili.com/video/BV1Ja4y1i7AF
- 以 Go 语言为例解释什么是伪共享以及如何解决 https://mp.weixin.qq.com/s/aM5jhfqWVu5FZDn6krV3tg
- Go中秘而不宣的数据结构 CacheLinePad：精细化优化“伪共享” https://colobu.com/2024/11/17/go-internal-ds-cacheline/


## 内存管理进程
核心：分层缓存机制
- 针对不同的内存场景采⽤不同的独特解决⽅式，提⾼局部性逻辑和细微粒度内存的复⽤率。

### 操作系统
- 虚拟内存
- MMU 内存管理单元（管理虚拟内存和物理内存映射关系）
- 虚拟内存本身怎么存放（页表 -> 页 -> PTE 页表条目）
- CPU内存访问过程
  - 如果虚拟页在主存中则页命中，反正为未命中（缺页），需要重新分配虚拟页，然后再走一遍寻址。
  - PS：命中率是访问次数与⻚命中次数之⽐。⼀般命中率低说明物理内存不⾜，数据在内存和磁盘之间交换频繁，但如果物理内存充分，则不会出现频繁的内存颠簸现象。
- 内存的局部性
  - 如果要设计出⼀个更加⾼效的程序，则提⾼代码的局部性访问是⾮常有效的程序性能优化⼿段之⼀

### TCMalloc 内存管理模型
三级缓存结构
- ThreadCache: 每个系统线程 M 都有对应的 ThreadCache
- CentralCache: 所有的 ThreadCache 共用一个 CentralCache
- PageHeap: 所有系统线程 M 和 CentralCache 都可以访问的全局资源，通过系统调⽤从操作系统虚拟内存中申请内存，为中对象和大对象设计的

TCMalloc 相关结构
- Page 将虚拟内存空间划分为多份同等⼤⼩的Page，每个Page默认为8KB（类似于操作系统的物理页）
- Span 多个连续的 Page（类似于操作系统的页表）
- Size Class 管理不同刻度的小对象 Span 链表集合（8B, 16B, 32B ...）

TCMalloc 的小对象分配（ - 256KB）
- 先 ThreadCache，从 FreeList 申请资源（不加锁）
- 再 CentralCache，从 CentralFreeList 申请资源（加锁）
- 最后 PageHeap（加锁）


TCMalloc 的中对象分配（256KB - 1MB，即128个Page）
- 直接向 PageHeap 发起申请 Span 请求
  - 如果属于 小 Span（即小于128个Page） 申请，则按刻度从对应的 Span 集合中获取指定长度的 Span 链表
  - 否则按大对象分配逻辑


TCMalloc 的大对象分配（ 1MB，即128个Page - ）
- 直接向 PageHeap 发起申请 Span 请求
  - 根据刻度从 Large Span Set 集合获取合适的 Span

### Go 基于 TCMalloc 实现的内存管理模型
三级缓存结构
- ThreadCache -> MCache
- CentralCache -> MCentral
- PageHeap -> MHeap


管理单元相关结构
- Page -> Page
- Span -> mSpan
- Size Class -> 细分为下面，管理 Span
  - Object Size
    - Page是Go语⾔内存管理与操作系统交互衡量内存容量的基本单元，Go语⾔内存管理内部⽤来给对象存储内存的基本单元是Object
  - Size Class
    - 于TCMalloc 概念一致，即刻度
  - Span Class
    - 额外定义，⼀个Size Class会对应两个Span Class，其中⼀个Span为存放需要GC扫描的对象（包含指针的对象），另⼀个Span为存放不需要GC扫描的对象（不包含指针的对象）

MCache（不加锁）
- MCache与Go语⾔协程调度模型GPM中的P绑定，⽽不是和线程绑定。因为Go语⾔调度的GPM模型，真正可运⾏的线程M的数 与P的数 ⼀致，即GOMAXPROCS个，所以MCache与P进⾏绑定更能节省内存空间，可以保证每个G使⽤MCache时不需要加锁就可以获取内存，⽽TCMalloc中的ThreadCache随着Thread的增多，ThreadCache的数 相对成正⽐增多。
- PS：申请 size=0 时的特殊处理，不做具体的分配，只返回一个固定的地址

MCentral（加锁）
- 协程逻辑层与MCache的内存交换单位是Object，MCache与MCentral的内存交换单位是Span，⽽MCentral与MHeap的内存交换单位是Page。
- MCentral与TCMalloc中的Central不同的是MCentral针对每个Span Class级别有两个Span链表，⽽TCMalloc中的Central只有⼀个。
- 属于 MHeap 的一部分

MHeap（加锁）
- MHeap是内存块的管理对象，通过Page对内存单元进⾏管理
- Arenas（由所有的 HeapArena 组成）
  - HeapArena（详细管理每⼀系列Page的结构）
    - ⼀个HeapArena占⽤内存64MB，其中⾥⾯的内存是⼀个⼀个的mspan，当然最⼩单元依然是Page
    - 每个HeapArean包含⼀个bitmap，其作⽤是标记当前这个HeapArena的内存使⽤情况。其主要服务于GC垃圾回收模块，bitmap共有两种标记，⼀种是标记对应地址中是否存在对象另⼀种是标记此对象是否被GC模块标记过，所以当前HeapArena中的所有Page均会被bitmap所标记。
    - ArenaHint为寻址HeapArena的结构，其有三个成员
      - (1)addr为指向的对应HeapArena的⾸地址。
      - (2)down为当前的HeapArena是否可以扩容。
      - (3)next指向下⼀个HeapArena所对应的ArenaHint的⾸地址。

Tiny 对象分配（ 1 - 16B）
- ⼤地使⽤微⼩对象时可能会对Size Class=1的Span造成浪费，所以Go语⾔内存管理决定尽 不使⽤Size Class=1的Span，⽽是将申请的Object⼩于16B的申请统⼀归类为Tiny对象申请。
- 步骤
   - P向MCache申请微⼩对象，如⼀个Bool变 。如果申请的Object在Tiny对象的⼤⼩范围，则进⼊Tiny对象申请流程
   - 判断申请的Tiny对象是否包含指针，如果包含指针，则进⼊⼩对象申请流程
   - 如果Tiny空间的16B没有多余的存储容 ，则从Size Class=2（Span Class=4或5）的Span中获取⼀个16B的Object放置于Tiny缓冲区
   - 将1B的Bool类型放置在16B的Tiny空间中，以字节对⻬的⽅式放置
   -

小对象分配（16B - 32KB）
- 向 MCache 申请分配
- 步骤
  - 计算大小，如果是小对象，进入分配流程
  - 根据Size匹配对应的Size Class内存规格，进行分配
  - 不够则向上 MCentral 或 MHeap 继续申请

大对象分配（ 32KB - ）
- 直接向 MHeap 发起申请分配
- 步骤
  - 从Arenas中的HeapArena申请相对应的Pages

