package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("data = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:4]
	printSlice(s)

	s = s[:2]
	printSlice(s)

	s = s[1:]
	printSlice(s)

	x := []int{99}
	printSlice(x)

	x = append(2)

}
