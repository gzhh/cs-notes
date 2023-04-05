package main

import (
	"fmt"
	"net/http"
)

func mayPanic() {
	panic("a problem")
}

// http server
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// demo1: catch error then exit（panic 在 defer recover 同一层级）
	mayPanic()
	fmt.Println("After mayPanic()")

	// demo2: catch error then stay（panic 在 defer recover 子层级）
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mayPanic()
		fmt.Println("After mayPanic()")
	})

	http.ListenAndServe(":8080", nil)
}
