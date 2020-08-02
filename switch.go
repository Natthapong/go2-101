package main

import (
	"fmt"
	"runtime"
)

func main () {
	switch os:= runtime.GOOS; os {
	case "linux":
		fmt.Println("it's linux") 
	case "darwin": 
		fmt.Println("it's darwin")
	default: 
		fmt.Println("default os ")
	}
}