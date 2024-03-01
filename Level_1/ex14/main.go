package main

import "fmt"

func checkType(unknown interface{}){
	switch unknown.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("float64")
	case []int:
		fmt.Println("[]int")
	case []string:
		fmt.Println("[]string")
	case []bool:
		fmt.Println("[]bool")
	case []float64:
		fmt.Println("[]float64")
	case nil:
		fmt.Println("nil")
	case chan struct{}:
		fmt.Println("chan")
	case func():
		fmt.Println("func")
	case struct{}:
		fmt.Println("struct")

	default:
		fmt.Println("unknown")
	}
}
func main(){
	var unknown interface{}

	unknown = "hello"
	checkType(unknown)
	unknown = 123
	checkType(unknown)
	unknown = true
	checkType(unknown)
	unknown = 123.45
	checkType(unknown)
	unknown = []int{1,2,3}
	checkType(unknown)
	unknown = []string{"hello","world"}
	checkType(unknown)
	unknown = []bool{true,false}
	checkType(unknown)
	unknown = []float64{1.2,3.4}
	checkType(unknown)
	unknown = nil
	checkType(unknown)
	unknown = make(chan struct{})
	checkType(unknown)
	unknown = func(){}
	checkType(unknown)
	unknown = interface{}(123)
	checkType(unknown)
	unknown = struct{}{}
	checkType(unknown)
}