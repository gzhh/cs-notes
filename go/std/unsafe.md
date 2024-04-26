# unsafe
- https://pkg.go.dev/unsafe
- https://pkg.go.dev/unsafe#Sizeof

## unsafe.Pointer
- https://pkg.go.dev/unsafe#Pointer

四种核心操作：
- 任何类型的指针值都可以转换为 unsafe.Pointer。
- unsafe.Pointer 可以转换为任何类型的指针值。
- uintptr 可以转换为 unsafe.Pointer。
- unsafe.Pointer 可以转换为 uintptr。

### 使用场景
- 1.unsafe.Pointer 可以让你的变量在不同的指针类型转来转去，也就是表示为任意可寻址的指针类型。
- 2.uintptr 常用于与 unsafe.Pointer 打配合，用于做指针运算，巧妙地很。

### 详细介绍
- 详解 Go 团队不建议用的 unsafe.Pointer https://eddycjy.com/posts/go/unsafe-pointer/
- https://go101.org/article/unsafe.html
