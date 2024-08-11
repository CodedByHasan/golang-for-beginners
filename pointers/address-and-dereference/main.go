package main

import "fmt"

func main() {
	i := 10

	// Address Operator (&) - Displays where i is stored in memory
	fmt.Printf("%T %v \n", &i, &i)

	// Derefernce Operator (*) - used to access the value at a particular memory address
	fmt.Printf("%T %v \n", *(&i), *(&i))
}
