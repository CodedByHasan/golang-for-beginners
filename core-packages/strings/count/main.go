package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "The quick brown fox jumped over the lazy dog."
	testString2 := "The quick brown fox jumped over The lazy dog."
	searchFor := "The"

	// Count function is case sensitive
	result1 := strings.Count(testString, searchFor)
	result2 := strings.Count(testString2, searchFor)

	fmt.Println(result1)
	fmt.Println(result2)
}
