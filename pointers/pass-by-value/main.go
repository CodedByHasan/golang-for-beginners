package main

import "fmt"

func modify(str string) {
	str = "world"
}

func main() {
	strA := "hello"
	fmt.Printf("BEFORE modify(): %s\n", strA)

	// Value remains unchanged
	// because strA is passed by value
	modify(strA)
	fmt.Printf("AFTER modify(): %s\n", strA)

}
