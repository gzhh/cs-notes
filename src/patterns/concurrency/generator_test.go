package concurrency

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	fmt.Println("No bottles of beer on the wall")

	for i := range Count(1, 99) {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
		// Pass it around, put one up, 1 bottles of beer on the wall
		// Pass it around, put one up, 2 bottles of beer on the wall
		// ...
		// Pass it around, put one up, 99 bottles of beer on the wall
	}

	fmt.Println(100, "bottles of beer on the wall")
}
