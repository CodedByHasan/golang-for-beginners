package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func calcAreaByVal(c Circle) {
	const PI float64 = 3.14
	var area float64

	area = (PI * c.radius * c.radius)
	c.area = area
}

func calcAreaByRef(c *Circle) {
	const PI float64 = 3.14
	var area float64

	area = (PI * c.radius * c.radius)
	(*c).area = area
}

func main() {
	c := Circle{
		x:      5,
		y:      5,
		radius: 5,
		area:   0,
	}

	fmt.Printf("Circle properties: %+v\n", c)

	// Pass by value - circle unchanged
	calcAreaByVal(c)
	fmt.Printf("Circle properties: %+v\n", c)

	// Pass by reference - properties changed
	calcAreaByRef(&c)
	fmt.Printf("Circle properties: %+v\n", c)

}
