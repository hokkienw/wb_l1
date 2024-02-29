package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}


func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.m[key]
}

func main() {
	sm := SafeMap{m: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go func(i int) {
			sm.Set(fmt.Sprintf("key%d", i), i)
			fmt.Println(sm.Get(fmt.Sprintf("key%d", i)))
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second)
}
