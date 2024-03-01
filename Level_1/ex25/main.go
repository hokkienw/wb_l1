package main

import (	
	"fmt"
	"time"
)

func Sleep(N time.Duration){
	<- time.After(N)
}

func main() {

fmt.Println("start")

Sleep(time.Second*5)

fmt.Println("end")

}