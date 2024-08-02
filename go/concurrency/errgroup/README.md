# errgroup
- https://pkg.go.dev/golang.org/x/sync/errgroup

Best Practice
- https://lailin.xyz/post/go-training-week3-errgroup.html


## Demo
```
package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// errgroup
func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.sogou.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			fmt.Println(url)

			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Failed fetched all URLs %+v\n", err)
	}

	time.Sleep(time.Second)
}
```
