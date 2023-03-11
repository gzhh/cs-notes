package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// http server 正确使用 defer recover
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// h0
	h0 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc h0!\n")
	}
	http.HandleFunc("/h0", h0)

	// h1
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		mayPanic()

		io.WriteString(w, "Hello from a HandleFunc h1!\n")
	}
	http.HandleFunc("/h1", h1)

	// h2
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		var jsonStr = `{"foo":"bar"}`
		var res map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &res)
		if err != nil {
			log.Fatal(err)
		}
		bar := res["foo"].(int)
		fmt.Println("foo", bar)

		io.WriteString(w, "Hello from a HandleFunc h2!\n")
	}
	http.HandleFunc("/h2", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
