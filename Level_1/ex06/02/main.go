package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Exiting.")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go worker(ctx)

    time.Sleep(5 * time.Second)

    cancel()

    time.Sleep(time.Second)
}
