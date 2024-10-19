# Garbage collector (GC)

## 介绍
- 一文搞懂七种基本的GC垃圾回收算法 https://mp.weixin.qq.com/s/M8R4QPidlCrr6vix4JUWmg
- 垃圾回收的算法与实现 https://book.douban.com/subject/26821357/

## 原理
实现原理
- A Guide to the Go Garbage Collector https://go.dev/doc/gc-guide
- Getting to Go: The Journey of Go's Garbage Collector https://go.dev/blog/ismmkeynote
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


