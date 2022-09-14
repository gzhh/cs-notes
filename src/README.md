## Golang notes
### builtin 相关
[iota](https://github.com/gzhh/golang-notes/tree/main/src/basic/builtin/iota.md)

### 内存相关
常见内存节省技巧
- 预先给切片分配容量
- 按类型调整结构体中字段顺序
- 使用 map[string]struct{} 代替 map[string]bool
