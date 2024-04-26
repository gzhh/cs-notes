package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func handle1(a int, b int) {
	defer recovery()
	//  The panic will not be recovered. This is because the recovery function is present in the different goroutine
	//  and the panic is happening in the divide() function in a different goroutine.
	go divide1(a, b)
}

func divide1(a int, b int) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func handle2(a int, b int) {
	//  The panic will be recovered.
	go divide2(a, b)
}

func divide2(a int, b int) {
	defer recovery()
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func main() {
	// handle1(5, 0)
	handle2(5, 0)
	fmt.Println("normally returned from main")
}
