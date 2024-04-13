# Generic
范型编程介绍
- https://en.wikipedia.org/wiki/Generic_programming
- https://zh.wikipedia.org/wiki/范型编程

## Go 中的范型编程
### 示例
- https://go.dev/doc/tutorial/generics
- https://gobyexample.com/generics

### 设计
- https://go.dev/blog/why-generics
- https://go.dev/blog/generics-next-step
- https://go.dev/blog/intro-generics
- https://go.dev/blog/when-generics

### 原理
- Golang 泛型初识 https://juejin.cn/post/7116815890343493646
- 简单易懂的 Go 泛型使用和实现原理介绍 https://zhuanlan.zhihu.com/p/509290914
- Go泛型是怎么实现的? https://colobu.com/2021/08/30/how-is-go-generic-implemented/
- Go泛型编程: interface 不再是那个interface https://colobu.com/2022/01/08/the-interface-is-not-that-interface-in-go-1-18/
- 泛型设计 https://golang3.eddycjy.com/posts/generics/
- Generics Are the Generics of Go https://www.capitalone.com/tech/software-engineering/generics-in-go/

### 第三方库
- https://github.com/samber/lo (lo - Iterate over slices, maps, channels...)

## Go 范型最佳实践
范型使用格式样例
```
// T is a type parameter.
func foo[T any](t T) {
    // ...
}
```

在数据结构中使用范型
```
// Uses a type parameter
type Node[T any] struct {
    Val  T
    next *Node[T]
}

// Instantiates a type receiver
func (n *Node[T]) Add(next *Node[T]) {
    n.next = next
}
```

范型适用的场景
- 数据结构，例如二叉树、链表、堆等
- 函数，与 any 类型的 slice, map, channel 等相关
- 抽象行为而不是类型，例如 sort package，可以使用范型进行抽象

范型不适用的场景
- 代码更复杂难以理解

## Usage
### 实现 getKeys 方法
```
package main

import "fmt"

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	m1 := make(map[string]int)
	m1["key1"] = 1
	m1["key2"] = 2
	m1["key3"] = 3
	fmt.Println(getKeys(m1))

	m2 := make(map[int]string)
	m2[1] = "value1"
	m2[2] = "value2"
	m2[3] = "value3"
	fmt.Println(getKeys(m2))

	m3 := make(map[[1]int]string)
	m3[[1]int{1}] = "value1"
	m3[[1]int{2}] = "value2"
	m3[[1]int{3}] = "value3"
	fmt.Println(getKeys(m3))
}
```

### 实现 getKeys 方法（自定义参数类型约束）
```
package main

import "fmt"

// Defines a custom type that restricts types to int and string
type customConstraint interface {
	~int | ~string
}

func getKeys[K customConstraint, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	m1 := make(map[string]int)
	m1["key1"] = 1
	m1["key2"] = 2
	m1["key3"] = 3
	fmt.Println(getKeys(m1))

	m2 := make(map[int]string)
	m2[1] = "value1"
	m2[2] = "value2"
	m2[3] = "value3"
	fmt.Println(getKeys(m2))

	m3 := make(map[[1]int]string)
	m3[[1]int{1}] = "value1"
	m3[[1]int{2}] = "value2"
	m3[[1]int{3}] = "value3"
	fmt.Println(getKeys(m3)) // can't compile
}
```
