package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a pipeline
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	errCh := make(chan error, 1) // Channel to capture errors

	// Start worker goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go func(id int) {
			defer wg.Done()
			processJobs(ctx, id, jobs, results)
		}(w)
	}

	// Producer function with error handling
	go func() {
		defer close(jobs)
		for j := 1; j <= 9; j++ {
			if j == 5 { // Simulating an error condition
				errCh <- fmt.Errorf("error encountered at job %d", j)
				cancel() // Cancel the context so workers stop processing
				return
			}
			select {
			case jobs <- j:
			case <-ctx.Done(): // Stop if context is canceled
				return
			}
		}
	}()

	// Error handling goroutine
	go func() {
		err := <-errCh
		fmt.Printf("Producer encountered an error: %v\n", err)
		cancel() // Cancel the context so all goroutines stop
	}()

	// Collector to close results after workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	totalResult := 0
	for r := range results {
		totalResult += r
		fmt.Printf("Got result: %d\n", r)
	}
	fmt.Printf("Final result: %d\n", totalResult)
}

func processJobs(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(500 * time.Millisecond) // Simulate work
			fmt.Printf("Worker %d processing job %d\n", id, job)
			results <- job * 2
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping due to cancellation\n", id)
			return
		}
	}
}
