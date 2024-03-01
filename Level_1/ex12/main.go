package main

import(
	"fmt"
)

func main(){
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	unique := make(map[string]struct{})
	for _, word := range sequence {
		unique[word] = struct{}{}
	}
	fmt.Println(unique)
}
