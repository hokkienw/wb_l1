package main

import (
	"fmt"
)

func setBitTo1(value int64, position uint) int64 {
	mask := int64(1) << position
	return value | mask
}

func setBitTo0(value int64, position uint) int64 {
	mask := int64(1) << position
	return value &^ mask
}

func main() {
	var number int64 = 42
	position := uint(2)
	fmt.Printf("Исходное число: %d\n", number)
	numberSetTo1 := setBitTo1(number, position)
	fmt.Printf("Установлен i-й бит в 1: %d\n", numberSetTo1)

	numberSetTo0 := setBitTo0(number, position)
	fmt.Printf("Установлен i-й бит в 0: %d\n", numberSetTo0)
}
