package main

import (
    "fmt"
    "time"
)

func worker(stopCh chan struct{}) {
    for {
        select {
        case <-stopCh:
            fmt.Println("Exiting.")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(time.Second)
        }
    }
}

func main() {
    stopCh := make(chan struct{})

    go worker(stopCh)

    time.Sleep(5 * time.Second)

    close(stopCh)

    time.Sleep(time.Second)
}