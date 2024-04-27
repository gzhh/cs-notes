# builtin 包
- https://pkg.go.dev/builtin
- https://github.com/golang/go/blob/go1.18/src/builtin/builtin.go

## builtin type
### 各类型大小分析
```
// 空结构体 struct{}
var s struct{}
fmt.Printf("sizeof(struct{}) = %d\n", unsafe.Sizeof(s))

// boolean
fmt.Printf("sizeof(bool) = %d\n", unsafe.Sizeof(true))

// float
var f32 float32 = 1.0
fmt.Printf("sizeof(float64) = %d\n", unsafe.Sizeof(f32))
var f64 float64 = 1.0
fmt.Printf("sizeof(float64) = %d\n", unsafe.Sizeof(f64))

// int and uint
var i int = 1
var ui uint = 1
var i32 int32 = 1
var ui32 uint32 = 1
fmt.Printf("sizeof(i) = %d\n", unsafe.Sizeof(i))
fmt.Printf("sizeof(ui) = %d\n", unsafe.Sizeof(ui))
fmt.Printf("sizeof(i32) = %d\n", unsafe.Sizeof(i32))
fmt.Printf("sizeof(ui32) = %d\n", unsafe.Sizeof(ui32))

// byte and rune
bytes := [4]byte{'1', '2', '3', '4'}
fmt.Printf("%+v\n", bytes)
fmt.Printf("sizeof(bytes) = %d\n", unsafe.Sizeof(bytes))

runes := [4]rune{'1', '2', '3', '4'}
fmt.Printf("%+v\n", runes)
fmt.Printf("sizeof(runes) = %d\n", unsafe.Sizeof(runes))

// pointer
fmt.Printf("sizeof(pointer) = %d\n", unsafe.Sizeof(&i)) // 指针为8个字节

// string
s := "1234"
for _, c := range s {
    fmt.Printf("%c ", c)
}
fmt.Printf("\nsizeof(string) = %d\n", unsafe.Sizeof(s)) // string 底层实现为包含一个指针和一个int类型长度的结构体

// any 或 interface{}
var a any = [0]byte{}
fmt.Printf("sizeof(any) = %d\n", unsafe.Sizeof(a)) // interface 底层实现为包含两个指针的结构体
```

What is the size of an int on a 64 bit machine?
- https://go.dev/doc/faq#q_int_sizes

uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
- uintptr

常量
- [iota](iota.md)

### 基础数据类型
- Go 指针和内存分配详解：https://segmentfault.com/a/1190000017473672
- Go 类型占用内存大小探究：https://chende.ren/2020/11/25172308-002-type-memory-size.html
- Go 的 []rune 和 []byte 区别：https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
- Go string、bytes、rune的区别：https://juejin.cn/post/6844903743524175879


## builting func
- append
- copy
- delete, clear, close
- panic, recover
- len, cap
- make, new
- max, min
- print, println


## builtin interface
- error
