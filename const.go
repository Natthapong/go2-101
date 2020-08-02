package main

import "fmt"

const Pi = 3.

const (
	_ = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	const world = "เวิร์ล"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
