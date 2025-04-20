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
- 万字长文讲透Go程序性能优化 (pprof & trace) https://mp.weixin.qq.com/s/wLPfiJ0wKH3DrBJS4yxeHw
- Go pprof性能分 https://zhuanlan.zhihu.com/p/492378357

pprof 和火焰图
- 如何阅读火焰图 https://www.kawabangga.com/posts/5861
- Flameshow 性能优化小记 https://www.kawabangga.com/posts/5427
- go pprof火焰图性能优化 https://zhuanlan.zhihu.com/p/71529062

gin pprof middleware
- https://github.com/gin-contrib/pprof

### 通过图片文件方式输出 Profile

安装 graphviz

- `brew install graphviz`

安装 go-torch

- `go get github.com/uber/go-torch`
- 下载并复制 [flamegraph.pl](http://flamegraph.pl) (https://github.com/brendangregg/FlameGraph) 至 $GOPATH/bin 路径下
- 将 $GOPATH/bin 加入 $PATH

安装 pprof

- `go get -u github.com/google/pprof`
- pprof 等价于 go tool pprof

分析

- go build xxx.go 编译go分析程序
- ./xxx 执行编译分析程序，生成分析文件 cpu.xxx、[mem.xxx](http://mem.xxx)、goutine.xxx
- 执行 pprof xxx [cpu.xxx](http://cpu.xxx) 分析
    - top 查看总的profile信息
    - list 查看具体函数的profile信息
    - ... 其他查看文档
    - svg 生成分析图片文件
- 执行 go-torch [cpu.xxx](http://cpu.xxx) 生成火焰图文件 torch.svg

### 通过 HTTP 方式输出 Profile

简单，适合于持续性运行的应用

1. 在应用程序中导入 import _ "net/http/pprof"，运行程序，然后浏览器访问 http://<host>:<port>/debug/pprof/ 即可展示 profiles 信息概览
2. 性能分析方式：当不断有用户请求时，使用下面两种方式分析程序性能
    - 命令行分析
        
        go tool pprof http://<host>:<port>/debug/pprof/profile?second=10（默认为30秒）
        
        - top
        - list
        - ...
    - 火焰图分析（生成 torch.svg）
        
        go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile（默认30秒）

### 问题分析
- Go heap profile 怎么了 https://colobu.com/2024/04/30/what-s-wrong-with-go-heap-profile/
- Can't use go tool pprof with an existing server https://stackoverflow.com/questions/30560859/cant-use-go-tool-pprof-with-an-existing-server


## Flame Graphs 火焰图
- https://www.brendangregg.com/flamegraphs.html
- https://github.com/brendangregg/FlameGraph
- Graphviz https://www.graphviz.org
- go-torch https://github.com/uber-archive/go-torch
- 6. Callgrind: a call-graph generating cache and branch prediction profiler https://valgrind.org/docs/manual/cl-manual.html


