package main

import (
	"fmt"
	"os"
)

func calcArea(radius float64) float64 {
	return 3.14 * radius * radius
}

func calcPerimeter(radius float64) float64 {
	return 2 * 3.14 * radius
}

func calcDiameter(radius float64) float64 {
	return 2 * radius
}

// Example of a higher-order function. This is a function that
// receives a function as an argument or returns a function as output.
func printResult(radius float64, calcFunction func(radius float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
}

func getFunction(query int) func(radius float64) float64 {
	queryToFunc := map[int]func(radius float64) float64{
		1: calcArea,
		2: calcDiameter,
		3: calcPerimeter,
	}
	return queryToFunc[query]
}

func checkQuery(query int) {
	if query != 1|2|3 {
		fmt.Println("Invalid query. Must be 1, 2 or 3.")
		os.Exit(1)
	}
}

func main() {
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f\n", &radius)

	fmt.Print("Enter \n 1) - Area \n 2) - Perimeter \n 3) - Diameter \n")
	fmt.Scanf("%d\n", &query)

	checkQuery(query)
	printResult(radius, getFunction(query))

}
