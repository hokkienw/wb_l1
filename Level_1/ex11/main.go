package main

import(
	"fmt"
)

func find(a []int, b []int){
	for _, v := range b{
		for _, j := range a{
			if v == j{
				fmt.Printf("%d ",j)
			}
		}
	}
	fmt.Println()
}

func find2(a []int, b []int) []int{
	c := []int{}
	for _, v := range b{
		for _, j := range a{
			if v == j{
				c = append(c, j)
			}
		}
	}
	fmt.Println(c)
	return c
}

func find3(a []int, b []int) map[int]struct{}{
	c:=  make(map[int]struct{})
	d:=  make(map[int]struct{})
	resault:= make(map[int]struct{})
	for _, v := range b{
		c[v] = struct{}{}
	}
	for _, v := range a{
		d[v] = struct{}{}
	}
	for k, _ := range c{
		if _, ok := d[k]; ok{
			resault[k] = struct{}{}
		}
	}
	fmt.Println(resault)
	return resault
}

func main(){
	a:= []int{1, 5 ,7, 3, 6, 9, 2, 4, 8}
	b:= []int{3, 8, 7, 4}
	find(a, b)
	find2(a, b)
	r:=find3(a, b)
	fmt.Println(r)
}
