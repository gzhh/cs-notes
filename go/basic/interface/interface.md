## interface 类型
### interface 的设计
interface 的使用
1. 尽可能的在 client 端定义 interface，在 service 端实现具体的方法，而不是直接在 service 端定义接口并实现所有的方法（类型 Java 的 OOP）。因为这样可以按 client 的需要进行依赖设计，减少无关方法的依赖。
2. 方法尽可能的接收 interface 参数，返回 struct，而不是返回 interface，因为这可能导致更复杂的抽象。除了在一些特殊场景，例如一些方法返回 io.Reader。


### interface 类型变量之间是否可比较分析
Compare whether two interfaces have identical dynamic types and equal dynamic values or if both are nil.

```
// only one empty interface type
type inf interface{}
var i1, i2 inf
fmt.Println(i1 == nil)
fmt.Println(i2 == nil)
fmt.Println(i1 == i2)
i1, i2 = 1, 1
fmt.Println(i1 == i2)
i1, i2 = 1, 2
fmt.Println(i1 == i2)
fmt.Println()

// two empty interfaces type
type inf3 interface{}
type inf4 interface{}
var i3 inf3
var i4 inf4
fmt.Println(i3 == nil)
fmt.Println(i4 == nil)
fmt.Println(i3 == i4)
i3, i4 = 1, 1
fmt.Println(i3 == i4)
i3, i4 = 1, 2
fmt.Println(i3 == i4)
fmt.Println()

// two same unempty interfaces type
type inf5 interface{ foo() }
type inf6 interface{ foo() }
var i5 inf5
var i6 inf6
fmt.Println(i5 == nil)
fmt.Println(i6 == nil)
fmt.Println(i5 == i6)
fmt.Println()

// two diffrent unempty interfaces type
type inf7 interface{ foo() }
type inf8 interface{ bar() }
var i7 inf7
var i8 inf8
fmt.Println(i7 == nil)
fmt.Println(i8 == nil)
fmt.Println(i7 == i8) // can't be compared, compile error
```

### 判断 struct 类型是否实现接口
因为 Go 里的接口实现不是侵入式，所以很多时候，我们使用某些接口时候并不知道是否实现了哪些接口。

样例
```
package main

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
}

/*
func (m *MyStruct) MyMethod() {
}
*/

func main() {
	// 编译时检查（使用声明方式）
	var _ MyInterface = new(MyStruct)

	// 编译时检查（类型转换方式）
	var _ MyInterface = (*MyStruct)(nil)
}
```

Ref
- https://golang.design/go-questions/interface/detect-impl/
- https://skyao.io/learning-go/feature/interface/check.html
- https://davidchan0519.github.io/2019/05/07/go-interface-implement/
- https://pkg.go.dev/reflect#example-TypeOf
- https://gfw.go101.org/article/reflection.html

### Ref
- https://github.com/golang/go/blob/go1.18/src/runtime/runtime2.go
- https://halfrost.com/go_interface/
- https://en.wikipedia.org/wiki/SOLID
