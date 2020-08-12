package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("data = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[2])
	s[2] = 123
	fmt.Println(s[2])

	s = append(s, 4, 5, 6)
	fmt.Println(s)

	data := []int{1, 2, 3, 4}

	result := append(s, data...)

	fmt.Println(result)

	result = append(s, data[0], data[1], data[2], data[3])

	fmt.Println(result)

	//var y []int
	y := make([]int, 5, 10)
	printSlice(y)
	if y == nil {
		fmt.Println("y is nil")
	} else {
		fmt.Println("y is ", y)
	}
}
