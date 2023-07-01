## Operating Systems 笔记

[https://zh.wikipedia.org/zh-cn/操作系统](https://zh.wikipedia.org/zh-cn/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F)
[GNU](https://www.gnu.org/software/software.html)

### 功能

操作系统位于底层硬件与用户之间，是两者沟通的桥梁。用户可以通过操作系统的用户界面，输入命令。操作系统则对命令进行解释，驱动硬件设备，实现用户要求。以现代标准而言，一个标准PC的操作系统应该提供以下的功能：
- 进程管理（提供操作系统调整多个进程的功能）
- 内存管理（管理进程对应的存储器空间，虚拟内存管理技术）
- 文件系统（管理磁盘的系统）
- 网络系统（提供计算机系统之间通讯的功能）
- 安全机制
- 用户界面
- 驱动程序

### 计算机功能结构
https://en.wikipedia.org/wiki/Computer_architecture
https://en.wikipedia.org/wiki/64-bit_computing#32-bit_vs_64-bit


### 字符编码
参考
- ASCII https://zh.wikipedia.org/wiki/ASCII
- Unicode https://zh.wikipedia.org/wiki/Unicode
- UTF-8 https://zh.wikipedia.org/wiki/UTF-8
- GB2312 https://zh.wikipedia.org/wiki/GB_2312
- GBK https://zh.wikipedia.org/wiki/汉字内码扩展规范

扩展
- MySQL 的 utf8 和 utf8mb4
- Unicode 是不是只有两个字节，为什么能表示超过 65536 个字符？：https://www.zhihu.com/question/22881537/answer/22947465
- Go 的 []rune 和 []byte 区别：https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
- Go string、bytes、rune的区别：https://juejin.cn/post/6844903743524175879

MacOS 查看系统默认编码
```
$ locale
输出类似于 zh_CN.UTF-8
```

### Terminal 终端相关知识
相关参考
- https://unix.stackexchange.com/questions/21147/what-are-pseudo-terminals-pty-tty
- https://man7.org/linux/man-pages/man4/tty.4.html
- https://man7.org/linux/man-pages/man7/pty.7.html
- https://jvns.ca/blog/2022/07/20/pseudoterminals/
- https://github.com/freman/goterm
- https://github.com/creack/pty
- https://github.com/chengjoey/web-terminal
