package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for i := range words {
		//fmt.Println(words[i])
		result[words[i]] = result[words[i]] + 1
	}

	return result
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Println(w)
}
