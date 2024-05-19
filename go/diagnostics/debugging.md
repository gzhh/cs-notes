# debugging

debug package
- https://pkg.go.dev/runtime/debug
- Functions
  - https://pkg.go.dev/runtime/debug#pkg-functions
  - https://pkg.go.dev/runtime/debug#SetTraceback

## Best Practice
- Some ways to get better at debugging https://jvns.ca/blog/2022/08/30/a-way-to-categorize-debugging-skills/
- Linux后台开发调试经验分享 https://mp.weixin.qq.com/s/Xk2-vfocgOm13sI1TTkZNw

开发中调试
- 尽量不要边开发边调试，例如打断点调试；而是开发完一个模块然后写单元测试进行调试。

Golang Troubleshooting 快速定位问题方法
- 1.链路追踪 tracer
- 2.sentry 日志
- 3.pprof 查看堆栈信息
- 4.GDB 或 Delve
  - GDB 或 Delve 都不行的话可以摘掉一个问题 pod，进到 pod 中通过 gdb 或者 dlv 打断点慢慢定位。


## vscode-go
Go for Visual Studio Code
- https://github.com/golang/vscode-go
- https://github.com/golang/vscode-go/wiki/debugging
- https://code.visualstudio.com/docs/editor/debugging
- https://code.visualstudio.com/docs/languages/go
- https://code.visualstudio.com/docs/languages/python
- https://code.visualstudio.com/docs/python/debugging

最佳实践
- vscode 调试技巧｜程序不是写出来的？是调出来的！https://mp.weixin.qq.com/s/6Br14gaZL1AO46KnivKuDA
- 【Vscode】调试go语言程序的最佳实践 https://cloud.tencent.com/developer/article/2028975
- How To Debug Go Code with Visual Studio Code https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
- Go Day 2021 on Google Open Source Live | Building Better Projects with the Go Editor https://www.youtube.com/watch?v=jMyzsp2E_0U


## GDB
GDB：The GNU Debugger
- https://www.sourceware.org/gdb/
- GDB中应该知道的几个调试方法 https://coolshell.cn/articles/3643.html
- 用GDB调试程序 https://blog.csdn.net/haoel/article/details/2879

GDB and Go
- Debugging Go Code with GDB https://go.dev/doc/gdb
- Go Web 编程 - 11.2. 使用 GDB 调试 https://learnku.com/docs/build-web-application-with-golang/112-debugging-with-gdb/3223

问题分析
- Debug Go program with GDB on macOS https://stackoverflow.com/questions/52534287/debug-go-program-with-gdb-on-macos


## LLDB
The LLDB Debugger
- https://lldb.llvm.org
- GDB to LLDB command map. https://lldb.llvm.org/use/map.html

LLDB and Go
- 使用 LLDB 调试 Go 程序 https://colobu.com/2018/03/12/Debugging-Go-Code-with-LLDB/


## Delve
Delve: 类似于 GDB 的 Debugger，但是只针对 Go 程序，功能比 GDB 更强大。
- https://github.com/go-delve/delve
- https://github.com/go-delve/delve/tree/master/Documentation

最佳实践
- 一个 Demo 学会使用 Go Delve 调试 https://mp.weixin.qq.com/s/Yz_p0S5N4ubf8wxLm5wbmQ
- 别再用GDB了，一文掌握Go最好用的调试器Delve https://mp.weixin.qq.com/s/Ed39t5I0k0ynfPch-ex7Ag
- go delve - The Golang Debugger https://earthly.dev/blog/golang-dlv/
- Debugging Go Code with Delve https://blog.devgenius.io/debugging-go-code-with-delve-22fe9de7e380


## Core dump
- https://en.wikipedia.org/wiki/Core_dump

介绍
- Go Wiki: CoreDumpDebugging https://go.dev/wiki/CoreDumpDebugging
- 如何保留 Go 程序崩溃现场 https://mp.weixin.qq.com/s/_0uatR-rrM7REhGndIbpOg
- 让golang程序生成coredump文件并进行调试 https://www.cnblogs.com/apocelipes/p/17536722.html

Segmentation fault
- https://en.wikipedia.org/wiki/Segmentation_fault
- What is a segmentation fault? https://stackoverflow.com/questions/2346806/what-is-a-segmentation-fault

### objdump
- https://en.wikipedia.org/wiki/Objdump

Objdump disassembles executable files.
- https://pkg.go.dev/cmd/objdump


