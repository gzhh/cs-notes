package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "巴塞罗那"
	textEncodedStr := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(textEncodedStr)

	data, err := base64.StdEncoding.DecodeString(textEncodedStr)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%s\n", data)
}
