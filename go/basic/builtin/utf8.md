## UTF-8
utf8 包主要是针对 UTF-8 字符编码文本的支持，提供了一些 runes 和 UTF-8 字节序列之间转换的函数。


## 常用函数
### Count Rune
```
package main

import (
	"fmt"
	"unicode/utf8"
)

// func RuneCount(p []byte) int
// func RuneCountInString(s string) (n int)
func main() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))
}
```

### Valid Rune
```
// func ValidRune(r rune) bool
func print1() {
	valid := 'a'
	invalid := rune(0xfffffff)

	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
}

// func Valid(p []byte) bool
// func ValidString(s string) bool
func print2() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}
```

### Decode Rune
```
// func DecodeRune(p []byte) (r rune, size int)
// func DecodeRuneInString(s string) (r rune, size int)
func main() {
	str := "Hello, 世界"

    // 多字节编码字符串遍历
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
}
```

### Encode Rune
```
// func EncodeRune(p []byte, r rune) int
func main() {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}
```

## 参考
- [字符编码笔记：ASCII，Unicode 和 UTF-8](http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)
- [UTF-8](https://en.wikipedia.org/wiki/UTF-8)
- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [utf8 package](https://pkg.go.dev/unicode/utf8)
