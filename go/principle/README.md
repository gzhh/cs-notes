# 底层原理

## Memory model and Race condition

**Memory model**

Ref: 
- [Memory model (programming)](https://en.wikipedia.org/wiki/Memory_model_(programming))
- [The Go Memory Model](https://go.dev/ref/mem)
- [Memory Models](https://research.swtch.com/mm)


**Race condition**

Ref:
- [Race condition](https://en.wikipedia.org/wiki/Race_condition)
- [Introducing the Go Race Detector](https://go.dev/blog/race-detector)
- [Data Race Detector](https://go.dev/doc/articles/race_detector)


**Malloc profilling**
- runtime/mprof.go
  - get goroutine id
    ```
    func getGid() uint {
        b := make([]byte, 64)
        b = b[:runtime.Stack(b, false)]
        fmt.Println(string(b))
        b = bytes.TrimPrefix(b, []byte("goroutine "))
        b = b[:bytes.IndexByte(b, ' ')]
        n, _ := strconv.ParseUint(string(b), 10, 64)
        return uint(n)
    }
    ```
