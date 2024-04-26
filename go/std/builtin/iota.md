## 自增常量 iota 用法

iota 是一个预定义定位符，表示无类型整数，初始化值为 0

通常 iota 需配合 const 关键字一起使用，多个常量使用 iota 赋值时，iota 值默认会按新行进行自增 1 操作（跳过空行和注释行）

### 不使用 iota 设置自增常量
```
const (
	Const0 = 0
	Const1 = 1
	Const2 = 2
```

### 使用 iota 常见用法
使用 iota，常量从 0 开始赋值，后续赋值依次加 1
```
const (
	Const0 = iota
	Const1
	Const2
)
// 0, 1, 2
```

设置常量初始值为 1，后续赋值依次加 1
```
const (
	Const0 = iota + 1
	Const1
	Const2
)
// 1, 2, 3
```

设置常量初始值为 1，后续赋值依次按前一个值的 2 倍进行赋值
```
const (
	Const0 = 1 << iota
	Const1
	Const2
)
// 1, 2, 4
```

### iota 更多用法
```
const (
	Const0 = iota
	Const1 = iota
	Const2 = iota
)
// 0, 1, 2
```

```
const (
	Const0 = iota
	Const1
	Const2 = iota
)
// 0, 1, 2
```

```
const (
	Const0 = iota

	Const1
	// comment
	Const2 = iota
)

// 0, 1, 2
// 自增逻辑会跳过空行和注释行
```

```
const (
	Const0 = iota
    _
	Const1
	Const2
)

// 0, 2, 3
// 使用占位符能跳过一次自增
```

```
const (
	Const0 = iota
	Const1
)
const (
	Const2 = iota
)

// 0, 1, 0
// 当 const 关键字重复使用时，iota 值会被重置
```

```
const (
	Const0 = iota
	Const1 = iota + 4
	Const2 = iota * 4
)
// 0, 5, 8
// iota 表达式，可以使用表达式给常量赋值
```
