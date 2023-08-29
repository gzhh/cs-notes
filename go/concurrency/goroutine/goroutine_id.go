package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 黑魔法：获取 goroutine id
func getGid() uint {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	fmt.Println(string(b))
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return uint(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i, getGid())
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(getGid())
}
