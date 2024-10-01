package main

import "fmt"

func calculate(a int, b int) []float64 {
	var results []float64 // results slice

	sum := a + b
	difference := a - b
	product := a * b
	quotient := float64(a) / float64(b)

	// appending to slice
	results = append(results, float64(sum), float64(difference), float64(product), quotient)

	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
