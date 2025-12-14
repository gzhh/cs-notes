# Basic

- [Serialization](serialization)
- [Configuration file](configuration-file)

## Character encoding
- https://en.wikipedia.org/wiki/Character_encoding
- https://zh.wikipedia.org/wiki/字符编码
- [一文读懂字符编码](https://mp.weixin.qq.com/s/5pAgcjk_lFGrPhSUp2Na3Q)
- https://en.wikipedia.org/wiki/ASCII
- https://en.wikipedia.org/wiki/Unicode
- https://en.wikipedia.org/wiki/UTF-8
- http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html


## 强/弱类型语言
强弱类型语言主要区别在于隐式类型转换的严格程度。 强类型语言不允许或严格限制不同数据类型间的隐式转换，例如在 Python 中 `1 + "2"` 会报错，需要显式转换。 弱类型语言则允许变量自动进行类型转换，例如在 JavaScript 中 `1 + "2"` 的结果是字符串 `"12"`。 

强类型语言
- 核心特点：: 不允许隐式类型转换，必须通过显式强制转换才能改变变量类型。 
- 优点：: 类型安全，能更早地暴露错误，代码可预测性强。 
- 缺点：: 代码灵活性相对较低，需要手动处理类型转换。 
- 代表语言：: Python、Java、Go、C#。 

弱类型语言
- 核心特点：: 允许变量在不同上下文或运算中自动进行类型转换。
- 优点：: 编写代码更灵活，简化操作。
- 缺点：: 容易导致难以预料的错误和bug，调试困难。
- 代表语言：: JavaScript、PHP、C（部分场景）。 


## 静态类型与动态类型语言
静态语言和动态语言的主要区别在于类型检查的时机：静态语言在编译时检查类型，而动态语言在运行时检查类型。静态语言通常具有更好的性能和安全性，适合开发大型系统，而动态语言则更灵活，适合快速原型开发。

静态语言
- 类型检查: 在编译阶段进行，数据类型在编译时即可确定。 
- 优势:
  - 性能: 通常比动态语言更快，因为代码在运行时无需进行类型检查。 
  - 安全性: 可以在编译时发现潜在的类型错误。 
- 典型语言: C, C++, Java, Go, C#, Rust。 

动态语言
- 类型检查: 在运行时进行，变量的类型在运行时根据其赋值来确定。 
- 优势:
  - 灵活性: 更灵活，可以在运行时改变程序结构和行为。 
  - 开发效率: 便于快速原型开发，可以更专注于业务逻辑的实现。 
- 典型语言: Python, JavaScript, Ruby, PHP, Lua

