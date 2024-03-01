package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	value int
}
func (c *Counter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main(){
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
		wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}