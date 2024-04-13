<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Golang notes](#golang-notes)
  - [标准库](#%E6%A0%87%E5%87%86%E5%BA%93)
  - [builtin 相关](#builtin-%E7%9B%B8%E5%85%B3)
  - [错误处理](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
  - [内存相关](#%E5%86%85%E5%AD%98%E7%9B%B8%E5%85%B3)
  - [time 相关](#time-%E7%9B%B8%E5%85%B3)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# The Go Programming Language
## Golang notes
### 标准库
- https://pkg.go.dev/std
- https://pkg.go.dev/runtime
- https://pkg.go.dev/sync

### builtin 相关
基础数据类型
- Go 指针和内存分配详解：https://segmentfault.com/a/1190000017473672
- Go 类型占用内存大小探究：https://chende.ren/2020/11/25172308-002-type-memory-size.html
- Go 的 []rune 和 []byte 区别：https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
- Go string、bytes、rune的区别：https://juejin.cn/post/6844903743524175879

常量
- [iota](https://github.com/gzhh/golang-notes/tree/main/src/basic/builtin/iota.md)

参考
- https://pkg.go.dev/builtin
- https://pkg.go.dev/unsafe#Sizeof
- https://pkg.go.dev/unsafe#Pointer


### [错误处理 Error handling](basic/error-handling)
参考
- https://pkg.go.dev/errors
- https://go.dev/blog/error-handling-and-go
- https://go.dev/blog/defer-panic-and-recover
- https://earthly.dev/blog/golang-errors/


### 内存相关
内存模型
- Memory Models: https://research.swtch.com/mm

常见内存节省技巧
- 预先给切片分配容量
- 按类型调整结构体中字段顺序
- 使用 map[string]struct{} 代替 map[string]bool


### time 相关
最佳实践：
- https://gobyexample.com/time
- https://gobyexample.com/time-formatting-parsing
- https://gobyexample.com/timeouts
- https://gobyexample.com/timers
- https://gobyexample.com/tickers

参考：
- https://pkg.go.dev/runtime
- https://pkg.go.dev/time
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/runtime/time.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/time.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/sleep.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/tick.go

