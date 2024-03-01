package main

import (
	"fmt"
	"sync"
	"flag"
)

func main() {
	N:= flag.Int("n", 5, "Number of workers")
	flag.Parse()

	dataChan := make(chan int)

	var wg sync.WaitGroup

	go func() {
		defer close(dataChan)
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()


	for i := 0; i < *N; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for data := range dataChan {
				fmt.Printf("Worker %d: %d\n", workerID, data)
			}
		}(i + 1)
	}


	wg.Wait()
}
