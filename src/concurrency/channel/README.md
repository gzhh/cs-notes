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
