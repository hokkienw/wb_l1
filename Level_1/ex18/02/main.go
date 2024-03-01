package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value atomic.Int64
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
	c.value.Add(1)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var counter Counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Inc(&wg)
	}
	wg.Wait()
	fmt.Println(counter.value.Load())
}
