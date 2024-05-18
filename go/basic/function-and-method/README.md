# 函数和方法
- define methods on types
- a method is just a function with a receiver argument.

## 函数传参
- https://go.dev/doc/faq#pass_by_value

As in all languages in the C family, everything in Go is passed by value. That is, a function always gets a copy of the thing being passed.

- 向一个函数传递一个 int 值，就会得到 int 的副本。
- 而传递一个指针值就会得到指针的副本，但不会得到它所指向的数据。
- map 和 slice 的行为类似于指针：它们是包含指向底层 map 或 slice 数据的指针的描述符。
  - 复制一个 map 或 slice 值并不会复制它所指向的数据。
  - 复制一个接口值会复制存储在接口值中的东西。
  - 如果接口值持有一个结构，复制接口值就会复制该结构。如果接口值持有一个指针，复制接口值会复制该指针，但同样不会复制它所指向的数据。

值传递，引用类型
- map
- slice
- chan

深入理解
- 又吵起来了，Go 是传值还是传引用？https://mp.weixin.qq.com/s/qsxvfiyZfRCtgTymO9LBZQ

## 函数返回
函数返回的顺序
- 计算函数的返回值
- 执行 defer 语句
- 执行 return 语句

例子
```
// 注意这里被 defer 函数内变量 i 的作用域
func c() (i int) {
    defer func() { i++ }()
    return 1
}
// return 2
```

参考
- Go 中 defer 和 return 执行的先后顺序 https://www.cnblogs.com/saryli/p/11371912.html
- Go函数调用惯例：多参数返回 https://mp.weixin.qq.com/s/f2Zu-sj-i4nO0wmmOBkoUg

## receiver type
### value and pointer receivers 是适用场景
- https://go.dev/doc/faq#methods_on_values_or_pointers

A receiver must be a pointer
- If the method needs to mutate the receiver. This rule is also valid if the receiver
- If the method receiver contains a field that cannot be copied: for example, a type part of the sync package (we will discuss this point in mistake #74, “Copy- ing a sync type”).

A receiver should be a pointer
- If the receiver is a large object. Using a pointer can make the call more effi- cient, as doing so prevents making an extensive copy. When in doubt about how large is large, benchmarking can be the solution; it’s pretty much impossible to state a specific size, because it depends on many factors.

A receiver must be a value
- If we have to enforce a receiver’s immutability.
- If the receiver is a map, function, or channel. Otherwise, a compilation error occurs.

A receiver should be a value
- If the receiver is a slice that doesn’t have to be mutated.
- If the receiver is a small array or struct that is naturally a value type without mutable fields, such as time.Time.
- If the receiver is a basic type such as int, float64, or string.


## Method expressions & Method values
- https://go.dev/ref/spec#Method_expressions
- https://go.dev/ref/spec#Method_values

map with callback function
- [demo](map_with_callback_function.go)

Ref
- https://stackoverflow.com/questions/39936099/syntax-for-map-with-callback-to-functions-which-have-a-receiver
- 
