package main

import "fmt"

func removeElement(nums *[]int, index int) {
	*nums = append((*nums)[:index], (*nums)[index+1:]...)
	
}


func main() {
	a:= []int{1,2,3,4,5,6,7,8,9}
	index:= 2
	removeElement(&a, index)
	fmt.Println(a)
}