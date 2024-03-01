package main

import(
	"fmt"
	 "math"
)



func main(){
	a:= [8] float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	temperatureGroups := make(map[int][]float64)

	for _, temp := range a {
		roundedTemp := int(math.Trunc(temp / 10.0) * 10.0)

		temperatureGroups[roundedTemp] = append(temperatureGroups[roundedTemp], temp)
	}

	for group, temps := range temperatureGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}

}