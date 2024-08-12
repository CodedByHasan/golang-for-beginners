package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

// Example of a method with receiver of type *Circle
func (c *Circle) calcAreaPassByRef() {
	c.area = 3.14 * c.radius * c.radius
}

// Another example of a method with receiver type of Circle
func (c Circle) calcAreaPassByVal() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	fmt.Printf("BEFORE calling any method: %+v\n", c)

	// Pass by value - c will be unchanged
	c.calcAreaPassByVal()
	fmt.Printf("AFTER calcAreaPassByVal() %+v\n", c)

	// Pass by ref - c will be changed
	c.calcAreaPassByRef()
	fmt.Printf("AFTER calcAreaPassByRef() %+v\n", c)
}
