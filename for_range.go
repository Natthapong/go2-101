package main

import "fmt"

func main() {
	names := []string{"A", "B", "C", "D", "E"}
	for index, value := range names {
		fmt.Println(index, value)
	}
}
