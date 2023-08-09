# Generic

https://go.dev/blog/intro-generics

https://go.dev/doc/tutorial/generics

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
