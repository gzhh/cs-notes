# Context
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
- https://pkg.go.dev/context
- https://pkg.go.dev/golang.org/x/net/context


## 介绍
- Go Concurrency Patterns: Context https://go.dev/blog/context
- Go 1.7 Context https://go.dev/doc/go1.7#context
- Contexts and structs https://go.dev/blog/context-and-structs


## 实践
- https://gobyexample.com/context
- https://golang.cafe/blog/golang-context-with-timeout-example.html


## 深入理解
- Go 并发编程实战课 - 11 | Context：信息穿透上下文 https://www.pursuitking.com/go_2/1-11.html
- The Complete Guide to Context in Golang: Efficient Concurrency Management https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea
- 【Go】透彻理解 Context https://mp.weixin.qq.com/s/BN1fUdwiz9QDZCsEmYjddA
- Go 组件 Context 源码学习笔记 https://mp.weixin.qq.com/s/PoXSEDHRyKCyjibFGS0wHw
- Context 是怎么在 Go 语言中发挥关键作用的 https://mp.weixin.qq.com/s/8hR-TNhJ_N5NfaGhHZAdEQ
- Go语言实战笔记（二十）| Go Context https://www.flysnow.org/2017/05/12/go-in-action-go-context
- Go 语言设计与实现 - 6.1 上下文 Context https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
- Go’s Context in Distributed Systems https://medium.com/@sogol.hedayatmanesh/gos-context-in-distributed-systems-ccdec75d5868

### 问题分析
- 面试官：Context携带数据是线程安全的吗？https://mp.weixin.qq.com/s/EQOLlJ2VSPd2-0Ey4rfZVw
- Go Context 到底放第一个参数传，还是放结构体里？https://mp.weixin.qq.com/s/DZl_lIueSHpdjfWJU9RrHg
