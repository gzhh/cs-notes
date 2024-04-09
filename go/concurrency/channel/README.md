<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Channel](#channel)
    - [介绍](#%E4%BB%8B%E7%BB%8D)
    - [注意点](#%E6%B3%A8%E6%84%8F%E7%82%B9)
    - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Channel
### 介绍
- Go's concurrency primitives - goroutines and channels - provide an elegant and distinct means of structuring concurrent software.
- Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
- Like maps and slices, channels must be created before use: ch := make(chan int)
- Channels are the pipes that connect concurrent goroutines.

### 原理
- https://go.dev/src/runtime/chan.go
- https://www.pursuitking.com/go_2/1-13.html
- https://halfrost.com/go_channel/
- https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/
- https://clavinjune.dev/blog/buffered-vs-unbuffered-channel-in-golang-dc97bf/

### Channel 之间比较
```
c1 := make(chan int, 2)
c2 := make(chan int, 2)
c3 := c1
fmt.Println(c1 == c2) // false
fmt.Println(c1 == c3) // true
```


### 注意点
1. 往已关闭的channel中发送消息，会发生panic
2. 从已读完消息的channel中读取消息，如果不判断，会读channel指定类型的零值

### 参考
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
