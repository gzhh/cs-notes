# Memory allocator

## 原理
Memory Allocation
- The GNU Allocator (The GNU C Library)
- https://www.gnu.org/software/libc/manual/html_node/The-GNU-Allocator.html
  - 聊聊C语言中的malloc申请内存的内部原理 https://mp.weixin.qq.com/s/7ZyCXUABL0Urso4VeaEdBQ
  - Glibc内存管理-ptmalloc2 https://www.cnblogs.com/mysky007/p/12349508.html
  - 百度工程师带你探秘C++内存管理（理论篇）https://mp.weixin.qq.com/s/mZAteEkyu0QeFhe1ZXevZA
  - 百度工程师带你探秘C++内存管理（ptmalloc篇）https://mp.weixin.qq.com/s/ObS65EKz1c3jooQx6KJ6uw
-
- google/tcmalloc
- https://github.com/google/tcmalloc
  - TCMalloc Overview https://google.github.io/tcmalloc/overview.html
  - TCMalloc : Thread-Caching Malloc https://goog-perftools.sourceforge.net/doc/tcmalloc.html
  - 图解 TCMalloc https://zhuanlan.zhihu.com/p/29216091
- 
- jemalloc/jemalloc
- https://github.com/jemalloc/jemalloc
- 对比
  - ptmalloc、tcmalloc与jemalloc对比分析 https://www.cyningsun.com/07-07-2018/memory-allocator-contrasts.html
  - Linux内存管理（五）：PTmalloc、TCmalloc、Jemalloc比较 https://blog.csdn.net/MOU_IT/article/details/118464093
  - 存优化总结:ptmalloc、tcmalloc和jemalloc https://zhuanlan.zhihu.com/p/497509956
  - Testing Memory Allocators: ptmalloc2 vs. tcmalloc vs. hoard vs. jemalloc https://news.ycombinator.com/item?id=17457699
  - c - What are the differences between (and reasons to choose) tcmalloc/jemalloc and memory pools? https://stackoverflow.com/questions/9866145/what-are-the-differences-between-and-reasons-to-choose-tcmalloc-jemalloc-and-m

Go Memory Allocation
- 图解Go语言内存分配 https://zhuanlan.zhihu.com/p/59125443
- Memory Allocations https://go101.org/optimizations/0.3-memory-allocations.html


src/runtime 源码
- Memory allocator. https://go.dev/src/runtime/malloc.go
- Implementation of (safe) user arenas. https://go.dev/src/runtime/arena.go
- Page allocator. https://go.dev/src/runtime/mpagealloc.go
- Page heap. https://go.dev/src/runtime/mheap.go
- Central free lists. https://go.dev/src/runtime/mcentral.go
- https://go.dev/src/runtime/mbitmap.go
- https://go.dev/src/runtime/mcache.go
- Memory statistics. https://go.dev/src/runtime/mstats.go
- Malloc profiling. https://go.dev/src/runtime/mprof.go


## 最佳实践

**Malloc profilling**
- runtime/mprof.go
  - get goroutine id
    ```
    func getGid() uint {
        b := make([]byte, 64)
        b = b[:runtime.Stack(b, false)]
        b = bytes.TrimPrefix(b, []byte("goroutine "))
        b = b[:bytes.IndexByte(b, ' ')]
        n, _ := strconv.ParseUint(string(b), 10, 64)
        return uint(n)
    }
    ```
