package main

import "fmt"

func main() {
	a:= 1
	b:= 2

	fmt.Println("до обмена:",a, b)

	a, b = b, a

	fmt.Println("параллельное присваивание:", a, b)

	a = a + b
    b = a - b
    a = a - b

	fmt.Println("вычитание сложение:", a, b)
}