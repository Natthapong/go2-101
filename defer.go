package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// f, err := os.Open("filename.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer func(x int) {
			fmt.Println(x)
		}(i)
	}

	fmt.Println("done")
}
