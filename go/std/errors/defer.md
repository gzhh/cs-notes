# defer

## 概念
1.defer 的执行顺序
- 多个 defer，先进后出（栈）

2.defer 和 return 先后
- return语句后面的表达式先执行，defer后面的语句后执行

3.有名函数返回值遇⻅defer的情况

```
package main

import "fmt"

func returnButDefer() (t int) { // t 初始化为 0，并且作用域为该函数全域
	defer func() {
		t = t * 10
	}()

	return 1
}

func main() {
	fmt.Println(returnButDefer())
}
```

4.defer 和 panic
- 当遇到panic时，会遍历本协程的defer链表，并执行defer。在执行defer的过程中，如果遇到recover，则停止panic，返回recover处继续往下执行。如果没有遇到recover，则遍历完本协程的defer链表后，向stderr抛出panic信息。
- 不捕获异常的情况：在panic之后的defer无法被触发，因为执行语句并没有将将其压栈

5.defer 中包含 panic
- panic仅有最后一个可以被recover捕获。

6.defer下的函数参数包含子函数
- 先执行最外层形参函数，再执行内层
