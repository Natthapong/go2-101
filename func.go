package main

import "fmt"

func add(x, y int) (int) {
	return x + y
}

func main() {
	r := add(42,13)
	fmt.Println(r)
}
