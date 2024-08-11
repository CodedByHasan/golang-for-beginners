package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	var circle Circle

	// accessing fields and initialising
	circle.x = 1
	circle.y = 1
	circle.radius = 1

	fmt.Printf("%+v\n", circle)
}
