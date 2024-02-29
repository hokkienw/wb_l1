package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        fmt.Println("Working...")
        time.Sleep(time.Second)
    }
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go worker(&wg)

    time.Sleep(5 * time.Second)

    wg.Done()

    wg.Wait()

	fmt.Println("Exiting.")
}
