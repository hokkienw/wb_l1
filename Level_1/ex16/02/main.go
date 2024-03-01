package main

import (
	"fmt"
	"sort"
)

func main() {
	a:= []int{1, 5 ,7, 3, 6, 9, 2, 4, 8}
	sort.Ints(a)
	fmt.Println(a)

	b:= []string{"a", "c", "b", "e", "d"}
	sort.Strings(b)
	fmt.Println(b)

	c:= []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	sort.Float64s(c)
	fmt.Println(c)

	d:= []rune{'a', 'b', 'c', 'd', 'e'}
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	fmt.Printf("%c", d)
}