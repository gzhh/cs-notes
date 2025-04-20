package main

import "fmt"

func main() {
	// 只能捕捉同一个 goroutine 中的 panic
	// 跨协程异常处理参考：https://taoshu.in/go/goroutine-panic.html
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic error: %v\n", r)
		}
	}()

	testPanic()
}

func testPanic() {
	panic("test")

	// panic: index out of range
	// arr := []int{1, 2}
	// fmt.Printf("i=%d\n", arr[2])

	// panic: integer divide by zero
	// var i int
	// div := 5 / i
	// fmt.Printf("div=%d\n", div)

	// panic: interface conversion, type assertion
	// m := map[string]interface{}{
	// 	"key":   "age",
	// 	"value": 28,
	// }
	// key := m["key"].(string)
	// value := m["value"].(string)
	// fmt.Printf("key=%v value=%v\n", key, value)

	// panic: invalid memory address or nil pointer dereference
	// var p *int
	// fmt.Println(*p)

	// panic: send on closed channel
	// ch := make(chan int)
	// close(ch)
	// ch <- 1
}
