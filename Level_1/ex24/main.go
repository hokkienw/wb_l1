package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
} 

func (p Point) Distance2(p2 Point) float64 {
	return math.Sqrt(math.Pow(p.x - p2.x, 2) + math.Pow(p.y - p2.y, 2))
}

func main() {
	newA := NewPoint(3, 4)
	newB := NewPoint(5, 6)
	fmt.Println(Distance(newA, newB))
	fmt.Println(newA.Distance2(newB))
}
