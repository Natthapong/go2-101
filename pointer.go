package main

import "fmt"

func main() {
	// i := 42
	// fmt.Println("address: ", &i)

	// var p *int

	// p = &i

	// fmt.Println("p value: ", p)
	// fmt.Println("p value inside address: ", *p)

	// *p = 55

	// fmt.Println("p value: ", p)
	// fmt.Println("p value inside address: ", *p)
	// fmt.Println("i value: ", i)

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	s := "KBTG"
	ps := &s
	fmt.Println(*ps)

}
