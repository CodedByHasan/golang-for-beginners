package main

import "fmt"

func main() {
	s := "hello"
	fmt.Printf("TYPE: %T VALUE: %v\n", s, s)

	ptrToS := &s
	fmt.Printf("TYPE: %T VALUE: %v\n", ptrToS, ptrToS)

	// Dereferencing ptrToS and changing value
	// This will change value of s.
	*ptrToS = "world"
	fmt.Printf("TYPE: %T VALUE: %v\n", s, s)
}
