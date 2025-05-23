# Optimization

Best Practice
- https://go101.org/optimizations/101.html
- Golang高性能编程实践 https://mp.weixin.qq.com/s/VMzhyySny60zABnxlzlVjQ
- 我是如何实现Go性能5倍提升的？https://mp.weixin.qq.com/s/SlPdSoMs1po1l19uaNMrIQ
- Go函数指针是如何让你的程序变慢的？https://mp.weixin.qq.com/s/bcmvPbWV7nBi7wIfr-MR8w
- 高德Go生态的服务稳定性建设｜性能优化的实战总结 https://mp.weixin.qq.com/s/UHaCLhiIyLYVrba-nEUONA
- 深入浅出 Go 性能优化：从原理到实践 https://mp.weixin.qq.com/s/BSWtifFJTIvaW4HNe603uw


## Escape analysis
概念：
- 编译器在编译的过程中就能够确定程序的预知和对该变量的声明周期的预知，编译器的这个分析过程就称之是逃逸分析。Go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(Escape Analysis)，当发现变量的作用域没有超出函数范围时，就可以放在栈上，反之则必须分配在堆。
- 一般在给一个引用类对象中的引用类成员进行赋值时可能出现逃逸现象。可以理解为，访问一个引用对象实际上是底层通过一个指针来间接地访问，但如果再访问里面的引用成员，则会有第二次间接访问，这样操作这部分对象时极大可能会出现逃逸的现象。
  - Go语言中的引用类型有func(函数类型)、interface(接口类型)、slice(切片类型)、 map(字典类型)、channel(管道类型)和∗(指针类型)等。

参考
- https://go.dev/wiki/CompilerOptimizations#escape-analysis
- https://tip.golang.org/src/cmd/compile/internal/escape/escape.go
- 通过实例理解 Go 逃逸分析 https://mp.weixin.qq.com/s/wNIX6LeHnBsl1UQuRtmC1g
- golang 逃逸分析详解 https://zhuanlan.zhihu.com/p/91559562

案例
- 如果变量是[]interface{}数据类型，则通过[]赋值必定会出现逃逸
- 如果变量是map[string]interface{}类型且尝试通过赋值，则必定会出现逃逸
- 如果map[interface{}]interface{}类型尝试通过赋值，则会导致key和value的赋值出现逃逸
- 如果变量是map[string][]string数据类型，则赋值会发生[]string逃逸
- 如果变量是[]∗int数据类型，则赋值的右值会发生逃逸现象
- 如果变量是func(∗int)函数类型，则进行函数赋值，会使传递的形参出现逃逸现象
- 如果变量是func([]string)函数类型，则进行[]string{"value"}赋值时会使传递的参数出现逃逸现象
- 如果变量是chan[]string数据类型，则向当前channel中传输[]string{"value"}时会发生逃逸现象

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
- 跟面试官聊 Goroutine 泄漏的 6 种方法，真刺激！ https://eddycjy.com/posts/go/goroutine-leak/
- Goroutine泄漏的危害、成因、检测与防治 https://juejin.cn/post/7128665615383920677
- Deep Dive into Go Memory Leak Debugging: A Practical Guide https://dev.to/jones_charles_ad50858dbc0/deep-dive-into-go-memory-leak-debugging-a-practical-guide-1ojm


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

