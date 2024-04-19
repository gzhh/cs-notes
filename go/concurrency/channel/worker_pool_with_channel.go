package main

import (
	"fmt"
	"sync"
	"time"
)

type Rows struct {
	ID int `json:"id"`
}

func handle(rows []*Rows) error {
	numWorkers := 10
	workerCh := make(chan *Rows, numWorkers)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		fmt.Println("Start worker ", i)
		wg.Add(1)
		go func(i int, workerCh chan *Rows) {
			defer wg.Done()

			for row := range workerCh {
				if row.ID%5 == 0 {
					continue
				}

				fmt.Println("Executing task", row.ID)

				// sleep for simulation
				time.Sleep(time.Second * 1)
			}

			fmt.Println("End worker", i)
		}(i, workerCh)
	}

	for _, row := range rows {
		workerCh <- row
	}
	close(workerCh)

	wg.Wait()

	return nil
}

func main() {
	size := 105
	rows := make([]*Rows, 0, size)
	for i := 0; i < size; i++ {
		rows = append(rows, &Rows{
			ID: i,
		})
	}
	handle(rows)
}
