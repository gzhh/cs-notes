## Golang Troubleshooting
快速定位问题方法

### 链路追踪 tracer

### sentry 日志

### pprof 查看堆栈信息

### GDB 或 Delve
GDB：The GNU Debugger

Delve: 类似于 GDB 的 Debugger，但是只针对 Go 程序，功能比 GDB 更强大。

前面都不行的话可以摘掉一个问题 pod，进到 pod 中通过 gdb 或者 dlv 打断点慢慢定位
