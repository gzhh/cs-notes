<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Channel](#channel)
    - [介绍](#%E4%BB%8B%E7%BB%8D)
    - [注意点](#%E6%B3%A8%E6%84%8F%E7%82%B9)
    - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Channel
介绍
``` text
Go's concurrency primitives - goroutines and channels - provide an elegant and distinct means of structuring concurrent software.
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
Like maps and slices, channels must be created before use: ch := make(chan int)
Channels are the pipes that connect concurrent goroutines.
```

## 原理
源码
- https://go.dev/src/runtime/chan.go

深入理解 channel 原理
- https://colobu.com/2016/04/14/Golang-Channels/
- https://www.pursuitking.com/go_2/1-13.html
- https://halfrost.com/go_channel/
- https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/
- https://i6448038.github.io/2019/04/11/go-channel/
- https://kingjcy.github.io/post/golang/go-channel/
- 【Golang】channel 数据结构 阻塞、非阻塞操作 多路select https://www.bilibili.com/video/BV1kh411n79h

深入理解 channel 特性
- https://clavinjune.dev/blog/buffered-vs-unbuffered-channel-in-golang-dc97bf/
- https://go101.org/article/channel-closing.html


## Channel 之间比较
```
c1 := make(chan int, 2)
c2 := make(chan int, 2)
c3 := c1
fmt.Println(c1 == c2) // false
fmt.Println(c1 == c3) // true
```


## 注意点
1. 往已关闭的channel中发送消息，会发生panic
2. 从已读完消息的channel中读取消息，如果不判断，会读channel指定类型的零值


## 最佳实践
- 从Go channel中批量读取数据 https://colobu.com/2023/12/23/batch-read-from-channels/


## 参考
- https://go.dev/tour/concurrency/2
- https://go.dev/tour/concurrency/3
- https://go.dev/tour/concurrency/4
- https://gobyexample.com/channels
- https://gobyexample.com/channel-buffering
- https://gobyexample.com/channel-synchronization
- https://gobyexample.com/channel-directions
- https://gobyexample.com/timeouts
- https://gobyexample.com/non-blocking-channel-operations
- https://gobyexample.com/closing-channels
- https://gobyexample.com/range-over-channels
