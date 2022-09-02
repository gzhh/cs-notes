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
Channels are the pipes that connect concurrent goroutines.
Go's concurrency primitives - goroutines and channels - provide an elegant and distinct means of structuring concurrent software.

### 注意点
1. 往已关闭的channel中发送消息，会发生panic
2. 从已读完消息的channel中读取消息，如果不判断，会读channel指定类型的零值

### 参考
https://blog.golang.org/codelab-share
https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/
https://clavinjune.dev/blog/buffered-vs-unbuffered-channel-in-golang-dc97bf/
