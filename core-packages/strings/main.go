package main

import (
	"fmt"
	"strings"
)

func main() {
	joker := "Why so serious?"
	batman := "Because I'm batman"

	learning := "learning standard library in Go"
	fun := "library in Go"

	// Contains searches for exact match
	result1 := strings.Contains(joker, batman)
	result2 := strings.Contains(learning, fun)

	fmt.Println(result1) // returns false
	fmt.Println(result2) // returns true
}
