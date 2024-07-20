# Assembler, Compiler and Interpreter
- Execution https://en.wikipedia.org/wiki/Execution_(computing)
- Assembler https://en.wikipedia.org/wiki/Assembly_language#Assembler
- Compiler https://en.wikipedia.org/wiki/Compiler
- Interpreter https://en.wikipedia.org/wiki/Interpreter_(computing)
- https://zh.wikipedia.org/wiki/机器语言

介绍
- 浅谈汇编器、编译器和解释器 https://zhuanlan.zhihu.com/p/70949843
- https://juejin.cn/s/编译器和汇编器的区别

## Assembler
- Assembler https://en.wikipedia.org/wiki/Assembly_language#Assembler
- x86 https://en.wikipedia.org/wiki/X86
- x86 assembly language https://en.wikipedia.org/wiki/X86_assembly_language
- x86 instruction listings https://en.wikipedia.org/wiki/X86_instruction_listings

汇编语言入门
- 汇编语言 https://book.douban.com/subject/35038473
- x86-64 Assembly Language Programming with Ubuntu http://www.egr.unlv.edu/~ed/x86.html

Go's Assembler
- A Quick Guide to Go's Assembler https://go.dev/doc/asm
- plan9 assembly 完全解析 https://github.com/cch123/golang-notes/blob/master/assembly.md
- 肝了一上午的Golang之Plan9入门 https://mp.weixin.qq.com/s/8wnMvROFQkVTKZ-qe4_eqw
- 从hello world谈起 https://zhuanlan.zhihu.com/p/630298292
- Go语言高级编程 第3章 Go 汇编语言 https://chai2010.cn/advanced-go-programming-book/ch3-asm/index.html
- 得到Go程序的汇编代码的方法 https://colobu.com/2018/12/29/get-assembly-output-for-go-programs/
- Go语言汇编快速指南 https://www.51cto.com/article/704916.html
- Golang 汇编入门知识总结 https://mp.weixin.qq.com/s/tN27osC6K0NM-Laj9MbtsA
- Go汇编优化入门 https://mzh.io/2018/05/Go汇编优化入门.pdf

### 从 hello world 谈起，定位 Go 程序真正的入口
参考
- Go 启动流程 https://github.com/cch123/golang-notes/blob/master/bootstrap.md
- Go 语言汇编入门 —— 从输出 HelloWorld 说起 https://juejin.cn/post/6844903929713524744

现有一段打印 Hello World 的 Go 代码，我们可以通过 gdb 来定位 Go 程序真正的入口。

```
package main

func main() {
	println("Hello, world!")
}
```

1.首先使用 `go build -gcflags "-N -l" -o hello hello.go` 编译程序

2.然后使用 `gdb hello` 进入 gdb

3.继续在 gdb 内执行命令 `info files` 得到程序的 Entry point，例如 `Entry point: 0x105aa40`

4.最后根据不同版本的 gdb 来定位真实的程序入口
- 老版本，例如 9.2，使用命令 `b *0x453b60` 即可得到真正的程序入口
    ```
    (gdb) b *0x105aa40
    Breakpoint 1 at 0x453b60: file /usr/local/go/src/runtime/rt0_linux_amd64.s, line 8.
    ```
- 新版本，例如 12.1，使用命令 `disassemble 0x105aa40` 即可得到真正的程序入口
    ```
    (gdb) disassemble 0x105aa40
    Dump of assembler code for function _rt0_amd64_darwin:
    ```
5.可知程序的入口分别为
- runtime/rt0_linux_amd64.s 的函数 `_rt0_amd64_linux`
- runtime/rt0_darwin_amd64.s 的函数 `_rt0_amd64_darwin`

6.以 linux 为例，调用链路为
- 首先调用 runtime/asm_amd64.s 的函数 `_rt0_amd64`
- 再调用 runtime/asm_amd64.s 的函数 `runtime·rt0_go`

7.在函数 runtime.rt0_go 中会做一些一些初始化操作，具体启动顺序是
```
// The bootstrap sequence is:
//
//	call osinit
//	call schedinit
//	make & queue new G
//	call runtime·mstart
//
// The new G calls runtime·main.
```
- CALL	runtime·osinit(SB) -> runtime/os_linux.go 的 `func osinit`
- CALL	runtime·schedinit(SB) -> runtime/proc.go 的 `func schedinit`
  - stackinit() -> runtime/stack.go 的 `stackinit()`
  - mallocinit() -> runtime/malloc.go 的 `mallocinit()`
  - gcinit() -> runtime/mgc.go 的 `gcinit()`
- CALL	runtime·newproc(SB) -> runtime/proc.go 的 `func newproc`
- CALL	runtime·mstart(SB) -> runtime/asm_amd64.s `TEXT runtime·mstart` -> runtime/proc.go 的 `func mstart0`
  - 引用 runtime/runtime2.go 的 g, m, p 结构
  - 调用 schedule(), One round of scheduler: find a runnable goroutine and execute it.


## Compiler
- Abstract syntax tree https://en.wikipedia.org/wiki/Abstract_syntax_tree
- 编译原理实战课 https://time.geekbang.org/column/intro/100052801
  - Go语言编译器 https://time.geekbang.org/column/article/266379
- 用Go语言自制编译器 https://book.douban.com/subject/35909089/
- How Golang Compiler Works? https://harendra21.medium.com/how-golang-compiler-works-6d11fcce02a9
- How to execute an object file
  - https://blog.cloudflare.com/how-to-execute-an-object-file-part-1
  - https://blog.cloudflare.com/how-to-execute-an-object-file-part-2
  - https://blog.cloudflare.com/how-to-execute-an-object-file-part-3
  - https://blog.cloudflare.com/how-to-execute-an-object-file-part-4

GCC
- https://gcc.gnu.org

LLVM
- https://zh.wikipedia.org/wiki/LLVM
- https://github.com/llvm/llvm-project
- https://llvm.org
- https://clang.llvm.org
- LLVM基本概念入门 https://zhuanlan.zhihu.com/p/140462815


## Interpreter
- 用Go语言自制解释器 https://book.douban.com/subject/35909085/

