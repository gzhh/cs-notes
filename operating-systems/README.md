## Operating Systems 笔记

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
