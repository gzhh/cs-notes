package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for true {
			fmt.Println(<-ch)
		}
	}()

	var nums = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		time.Sleep(1 * time.Second)
		ch <- nums[i]
	}
}
