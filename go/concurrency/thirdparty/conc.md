# conc
- https://github.com/sourcegraph/conc


## conc iter
```
package main

import (
	"fmt"

	"github.com/sourcegraph/conc/iter"
)

func handle(elem *int) {
	println(*elem)
}

var f = func(num *int) int {
	return *num * *num
}

func main() {
	values := []int{1, 2, 3, 4, 5}

	iter.ForEach(values, handle)

	returns := iter.Map(values, f)
	fmt.Println(returns)
}
```

## conc stream
```
package main

import (
	"fmt"
	"time"

	"github.com/sourcegraph/conc/stream"
)

var f = func(num int) int {
	return num * num
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go mapStream(in, out, f)

	go func() {
		for el := range out {
			fmt.Println(el)
		}
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	i := 0
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			i++
			in <- i
		}
	}
}

func mapStream(
	in chan int,
	out chan int,
	f func(int) int,
) {
	s := stream.New().WithMaxGoroutines(10)
	for elem := range in {
		elem := elem
		s.Go(func() stream.Callback {
			res := f(elem)
			return func() { out <- res }
		})
	}
	s.Wait()
}
```
