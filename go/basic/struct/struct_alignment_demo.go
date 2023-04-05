package main

import (
	"fmt"
	"unsafe"
)

type X struct {
	a int8  // int8 1字节
	b int64 // int64 8字节
	c int32 // int32 4字节
	d int16 // int16 2字节
}

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
// a               b                     c           d
// 0%1=0          8%8=0                16%4=0      20%2=0  最后再按最大字节8进行补充对齐（补上22 23）

type Y struct {
	a int8
	d int16
	c int32
	b int64
}

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
// a   d   c       b
//   2%2=0       8%8=0

func main() {
	// 内存对齐
	x := X{}
	y := Y{}
	// 24 16
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(y))
}
