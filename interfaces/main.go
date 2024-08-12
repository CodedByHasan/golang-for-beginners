package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.breadth)
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	// Both rectangle and square implement
	// shape interface
	r := rectangle{length: 3, breadth: 4}
	s := square{side: 5}

	fmt.Println("RECTANGLE")
	printData(r)

	fmt.Println("SQUARE")
	printData(s)
}
