# Optimization

Best Practice
- https://go101.org/optimizations/101.html
- Golang高性能编程实践 https://mp.weixin.qq.com/s/VMzhyySny60zABnxlzlVjQ
- 我是如何实现Go性能5倍提升的？https://mp.weixin.qq.com/s/SlPdSoMs1po1l19uaNMrIQ
- Go函数指针是如何让你的程序变慢的？https://mp.weixin.qq.com/s/bcmvPbWV7nBi7wIfr-MR8w
- 高德Go生态的服务稳定性建设｜性能优化的实战总结 https://mp.weixin.qq.com/s/UHaCLhiIyLYVrba-nEUONA


## Escape analysis
- https://go.dev/wiki/CompilerOptimizations#escape-analysis
- https://tip.golang.org/src/cmd/compile/internal/escape/escape.go
- 通过实例理解 Go 逃逸分析 https://mp.weixin.qq.com/s/wNIX6LeHnBsl1UQuRtmC1g
- golang 逃逸分析详解 https://zhuanlan.zhihu.com/p/91559562

变量的生命周期
- 包级变量：存在于整个程序的运行过程，程序终止，包变量才会被回收。
- 局部变量：动态生命周期，从每次生成新变量语句执行的时候开始，变量一直存在，直到这个变量成为不可达时（unreachable)，在那时他的存储空间可能被回收
- 函数参数和返回值也是局部变量：每次函数被调用的时候被创建
- 编译器决定储存空间堆栈分配

go build 时可以使用 `--gcflags` 指定编译选项，可以通过如下命令查看可选参数

`go tool compile --help`

go build 时可以用如下命令查看是否发生内存逃逸

`go build -gcflags "-m -m" test.go`

`go build -gcflags "-m -l" test.go`

### Third Party
goleak - Goroutine leak detector
- https://github.com/uber-go/goleak
- 跟面试官聊 Goroutine 泄露的 6 种方法，真刺激！ https://eddycjy.com/posts/go/goroutine-leak/
- Goroutine泄露的危害、成因、检测与防治 https://juejin.cn/post/7128665615383920677


## PGO
- Profile-guided optimization preview https://go.dev/blog/pgo-preview
- Profile-guided optimization in Go 1.21 https://go.dev/blog/pgo
- PGO: 为你的Go程序提效5% https://colobu.com/2023/09/13/pgo/
- Cloudflare最佳实践：如何通过 Go PGO 回收 CPU
  - https://blog.cloudflare.com/reclaiming-cpu-for-free-with-pgo
  - https://mp.weixin.qq.com/s/yrMkH5xOLUS8O9V2H8nfmA

Doc
- Profile-guided optimization https://go.dev/doc/pgo
- Go Wiki: PGO Tools https://go.dev/wiki/PGO-Tools


### 深入理解
- Go和C++通用性能优化黑魔法——PGO！https://mp.weixin.qq.com/s/VhrgXnntsMCk9Ajw00clkQ
- Go1.21中的PGO技术详解 https://mp.weixin.qq.com/s/FJgAa8A2ZZWj1619yioyvg
- Golang 程序性能优化利器 PGO 详解
  - https://zhuanlan.zhihu.com/p/649642714
  - https://zhuanlan.zhihu.com/p/649742674
  - https://zhuanlan.zhihu.com/p/650027126

