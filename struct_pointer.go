package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	v = Vertex{}
	v = Vertex{X: 1}
	v = Vertex{1, 2, 3}

	//p := &v
	p := &Vertex{1, 2, 3}

	//v.X =15
	(*p).X = 19
	p.X = 19

	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", v.X)
	fmt.Printf("%#v\n", (*p).X)
	fmt.Printf("%#v\n", p.X)
}
