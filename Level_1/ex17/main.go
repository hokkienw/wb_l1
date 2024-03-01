package main

import (
	"fmt"
	"sort"
)
func binarySearch(a []int, target int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main(){
a:= []int{1,5,8,9,4}
sort.Ints(a)
target:= 1

f:= binarySearch(a, target)

fmt.Println(f)
}