# Basic types

### 引用类型
- func(函数类型)
- interface(接口类型)
- slice(切片类型)
- map(字典类型)
- channel(管道类型)
- ∗(指针类型)


### 基础类型

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

### 零值

- `0` for numeric types
- `false` for the boolean type
- `""` (the empty string) for strings

### 类型判断

```

// 判断数据类型: 方法1 .(type)
// .(type) 必须和 switch 搭配使用，前提是需要列举类型
func DataType(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case interface{}:
		return "interface"
  // TODO ...
	default:
		return "unknow"
	}
}
var i int = 123
fmt.Println(DataType(i))
var str string = "af"
fmt.Println(DataType(str))

// 判断数据类型: 方法2 反射reflect
var i int = 123
fmt.Println(reflect.TypeOf(i))
var str string = "af"
fmt.Println(reflect.TypeOf(str))
```

### 类型转换

显式转换

The expression T(v) converts the value v to the type T.

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

隐式转换

1. 依赖赋值符号右侧数据值的类型，与类型之相同

```
var i int
j := i // j is an int
```

2. 依赖赋值符号右侧数据常量值的精度，numeric constant

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

Type Assertion 类型断言

- A type assertion provides access to an interface value's underlying concrete value.

```
// NOTE: value i should be interface type.
t := i.(T)

t, ok := i.(T)
```

```
// 数据类型转换，interface{} => int
var i1 interface{} = 123
if i2, ok := i1.(int); ok {
	fmt.Println(i2)
} else {
	// ok == false
}

// 数据类型转换，interface{} => string
var i1 interface{} = "abc"
if i2, ok := i1.(string); ok {
	fmt.Println(i2)
} else {
	// ok == false
}
```

### 常量

Constants are declared like variables, but with the `const` keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the `:=` syntax.

```
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```

## 最佳实践
### 1. 整数常量INT_MAX INT_MIN最大值最小值

**无符号整型uint**其最小值是0，其二进制表示的所有位都为0

`const` `UINT_MIN uint = 0`

其最大值的二进制表示的所有位都为1

`const` `UINT_MAX = ^uint(0)`

**有符号整型int**

根据补码，其最大值二进制表示，首位0，其余1

`const` `INT_MAX = int(^uint(0) >> 1)`

根据补码，其最小值二进制表示，首位1，其余0

`const` `INT_MIN = ^INT_MAX`
