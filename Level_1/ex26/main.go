package main

import (
	"fmt"
	"strings"
)

func IsAllUnique(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func IsAllUnique2(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]bool)
	for _, r := range s {
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}

func main() {
s:= "abcd"
fmt.Println(IsAllUnique2(s))
fmt.Println(IsAllUnique(s))
s = "abCdefAaf"
fmt.Println(IsAllUnique2(s))
fmt.Println(IsAllUnique(s))
}