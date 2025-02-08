package main

import "time"

func main() {
	// derfer recover
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in f", r)
		}
	}()

	/*
		go func() {
			// can not recovered
			panic("Panic in goroutine")
		}()
	*/

	// can recovered (defer recover func can only handle same goroutine panic)
	panic("Panic")

	time.Sleep(time.second)
}
