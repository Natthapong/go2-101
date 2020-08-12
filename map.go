package main

import (
	"fmt"
)

type Student struct {
	id    int
	name  string
	class string
}
type Vertex struct {
	Lat, Long float64
}

var y = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.399967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var student = map[int]Student{
	1: {id: 1, name: "Natthapong", class: "A"},
	2: {id: 2, name: "Natchanon"},
	3: {id: 3, name: "Pornnapa"},
}

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["nat"] = "Natthapong"
	//fmt.Println(m["nat"])

	//fmt.Println(y)
	//fmt.Println(student)

	y := map[int]Student{}
	y[1] = Student{1, "Natthapong", "A"}
	y[2] = Student{2, "Natchanon", "B"}

	s, ok := student[1]
	if ok {
		fmt.Printf("%#v, %#v\n", s, ok)
	} else {
		fmt.Printf("Not found\n")
	}
	s, ok = student[6]
	if ok {
		fmt.Printf("%#v, %#v\n", s, ok)
	} else {
		fmt.Printf("Not found\n")
	}

}
