package main

import "fmt"

type Human struct {
	Name string
}

type Action struct {
	Human
}

func (h *Human) Hello() {
	fmt.Println("Hello! My name is", h.Name)
}

func (a *Action) Goodbye() {
	fmt.Println("Goodbye!")
}

func main(){
	action := Action{ Human: Human{ Name: "Rashit"}}

action.Hello()

action.Goodbye()


}