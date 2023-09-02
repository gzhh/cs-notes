# 函数和方法

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

## receiver type
### value and pointer receivers 是适用场景

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

