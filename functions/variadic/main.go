package main

import "fmt"

// This is a variadic function which takes
// in zero or more arguments of the same type
func sumNumbers(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Printf("Hey %s. Here are your subjects:\n", student)

	for num, subject := range subjects {
		fmt.Printf("%d. %s\n", num+1, subject)
	}
}

func main() {
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(1, 2))
	fmt.Println(sumNumbers(1, 2, 3))

	printDetails("Harry Potter", "Defence Against the Dark Arts", "Potions", "Botany")
}
