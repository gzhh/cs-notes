# pprof

Profiling Go Programs
- https://go.dev/blog/pprof

pprof package
- https://pkg.go.dev/runtime/pprof
- https://pkg.go.dev/net/http/pprof
- https://github.com/felixge/fgprof

pprof command
- https://pkg.go.dev/cmd/pprof

pprof tool
- https://github.com/google/pprof
- https://github.com/google/pprof/blob/main/doc/README.md


## 介绍

### Profile Descriptions
- allocs: A sampling of all past memory allocations
- block: Stack traces that led to blocking on synchronization primitives
- cmdline: The command line invocation of the current program
- goroutine: Stack traces of all current goroutines
- heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
- mutex: Stack traces of holders of contended mutexes
- profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
- threadcreate: Stack traces that led to the creation of new OS threads
- trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.


## 最佳实践
- Go 语言编程之旅 - 6.1 Go 大杀器之性能剖析 PProf（上）https://golang2.eddycjy.com/posts/ch6/01-pprof-1/
- golang 性能优化分析工具 pprof (上) - 基础使用介绍 https://www.cnblogs.com/jiujuan/p/14588185.html
- Profiling Go programs with pprof https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
- 你不知道的 Go 之 pprof https://darjun.github.io/2021/06/09/youdontknowgo/pprof/
- Go 中的性能分析和执行跟踪 https://mp.weixin.qq.com/s/7_b1NnoLxyE-kp7mnfXj4w
- 通过pprof定位groutine泄漏 https://mp.weixin.qq.com/s/UcOwzNHqhPE0ZjKR1maWBw
- 修复go tool pprof存在的“bug” https://mp.weixin.qq.com/s/WSfDnlXzX36NQNs1t3M-Bg
- 生产环境Go程序内存泄露，用pprof如何快速定位 https://mp.weixin.qq.com/s/PEpvCqpi9TPhVuPdn3nyAg
- pprof - 在现网场景怎么用 https://www.cnblogs.com/yjf512/p/18120513
- 从真实事故出发：golang 内存问题排查指北 https://mp.weixin.qq.com/s/HdSIC93HMbqvbQisCr186Q

pprof 和火焰图
- 如何阅读火焰图 https://www.kawabangga.com/posts/5861
- Flameshow 性能优化小记 https://www.kawabangga.com/posts/5427
- go pprof火焰图性能优化 https://zhuanlan.zhihu.com/p/71529062

gin pprof middleware
- https://github.com/gin-contrib/pprof

### 问题分析
- Go heap profile 怎么了 https://colobu.com/2024/04/30/what-s-wrong-with-go-heap-profile/
- Can't use go tool pprof with an existing server https://stackoverflow.com/questions/30560859/cant-use-go-tool-pprof-with-an-existing-server


## Flame Graphs 火焰图
- https://www.brendangregg.com/flamegraphs.html
- https://github.com/brendangregg/FlameGraph
- Graphviz https://www.graphviz.org
- go-torch https://github.com/uber-archive/go-torch
- 6. Callgrind: a call-graph generating cache and branch prediction profiler https://valgrind.org/docs/manual/cl-manual.html
