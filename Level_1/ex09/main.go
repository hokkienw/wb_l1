package main

import (
	"fmt"
	"sync"
)

func main() {
	inputChan := make(chan int)
	outputChan := make(chan int)

	var wg sync.WaitGroup

	go func() {
		defer close(inputChan)
		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			inputChan <- num
		}
	}()

	go func() {
		defer close(outputChan)
		for num := range inputChan {
			outputChan <- num * 2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range outputChan {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
