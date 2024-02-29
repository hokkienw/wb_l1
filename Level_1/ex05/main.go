package main

import(
	"fmt"
	"time"
	"flag"
)

func sender(ch chan int){
	for {
		ch <- 1
		time.Sleep(time.Second)
	}
}

func receiver(ch chan int){
	for {
		fmt.Println(<-ch)
	}
}

func main(){
	N:= flag.Int("n", 5, "number of seconds")

	flag.Parse()

	ch:= make(chan int)

	go sender(ch)

	go receiver(ch)

	time.Sleep(time.Second * time.Duration(*N))

	close(ch)
	// time.Sleep(time.Second)
}