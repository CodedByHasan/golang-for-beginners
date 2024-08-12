package main

import "fmt"

func main() {
	// example of a anonymous function
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("Type of x: %T\n", x)
	fmt.Println(x(20, 30))

	y := func(l int, b int) int {
		return l + b
	}(20, 30)

	fmt.Printf("Type of y: %T\n", y)
	fmt.Println(y)

}
