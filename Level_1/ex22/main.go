package main

import (
	"fmt"
	"math/big"
)

func main() {
	a:= new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	b:= new(big.Int).Exp(big.NewInt(3), big.NewInt(30), nil)
	fmt.Println("first num: ",a)
	fmt.Println("second num: ", b)

	result := new(big.Int).Add(a, b)
	fmt.Println("add: ", result)
	result = result.Sub(a, b)
	fmt.Println("sub: ", result)
	result = result.Mul(a, b)
	fmt.Println("mul: ", result)
	result = result.Div(b, a)
	fmt.Println("div: ", result) 
}
