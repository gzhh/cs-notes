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
