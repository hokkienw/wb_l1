package main

import "fmt"

func Square(squares chan int, arr []int){
	for _, number := range arr {
		squares <- number*number
	}
	close(squares)
}

func main() {
	arr:= []int{2,4,6,8,10}
	squares:= make(chan int)

	go Square(squares, arr)

	var resault int
	for sq:= range(squares){
		resault += sq
	}
	fmt.Println(resault)
}