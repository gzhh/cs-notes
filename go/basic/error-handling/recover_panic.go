package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	//  The panic will not be recovered. This is because the recovery function is present in the different goroutine
	//  and the panic is happening in the divide() function in a different goroutine.
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func main() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}
