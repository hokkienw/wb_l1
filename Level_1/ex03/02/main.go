package main

import (
	"fmt"
	"sync"
)

func main() {
	arr:= []int{2,4,6,8,10}
	var resault int

	var wg sync.WaitGroup

	for _, number := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			resault += num*num
		}(number)
		}
	wg.Wait()

	fmt.Println(resault)
}