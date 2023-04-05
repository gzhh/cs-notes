package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
	test2()
}

// output: 21
func test1() {
	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Println(x)
}

// three three three
func test2() {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

/**
output: 132
等价于
s.Add(1)
defer s.Add(2)
s.Add(3)
*/
func test3() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}
