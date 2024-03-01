package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
  
  return strings.Repeat("a", int(size))
}


func someFunc() {
  v := createHugeString(1 << 10)
  copyData := make([]byte, 100)
  copy(copyData, []byte(v[:100]))
  justString = string(copyData)
}

func main() {
  someFunc()
  fmt.Println(justString)
}
