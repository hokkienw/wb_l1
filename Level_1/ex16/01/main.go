package main

import "fmt"

func partition(arr []int, left, right int) int{
	pivot := arr[left]
	for left < right{
		for left < right && arr[right] >= pivot{
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot{
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}


func quickSort(arr []int, left, right int){
	if left < right{
		mid := partition(arr, left, right)
		quickSort(arr, left, mid-1)
		quickSort(arr, mid+1, right)
	}
}

func main(){
	arr := []int{5,4,3,2,1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}