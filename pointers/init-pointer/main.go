package main

import "fmt"

func main() {
	s := "hello"

	// Initialising a ptr to s using different methods
	var ptrA *string = &s
	fmt.Println(ptrA)

	var ptrB = &s
	fmt.Println(ptrB)

	ptrC := &s
	fmt.Println(ptrC)

}
