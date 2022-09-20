## Golang notes
### 标准库
- https://pkg.go.dev/std
- https://pkg.go.dev/runtime
- https://pkg.go.dev/sync

### builtin 相关
基础数据类型
- Go 指针和内存分配详解：https://segmentfault.com/a/1190000017473672
- Go 类型占用内存大小探究：https://chende.ren/2020/11/25172308-002-type-memory-size.html
- Go 的 []rune 和 []byte 区别：https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
- Go string、bytes、rune的区别：https://juejin.cn/post/6844903743524175879

常量
- [iota](https://github.com/gzhh/golang-notes/tree/main/src/basic/builtin/iota.md)

参考
- https://pkg.go.dev/builtin
- https://pkg.go.dev/unsafe#Sizeof
- https://pkg.go.dev/unsafe#Pointer


### 内存相关
常见内存节省技巧
- 预先给切片分配容量
- 按类型调整结构体中字段顺序
- 使用 map[string]struct{} 代替 map[string]bool

