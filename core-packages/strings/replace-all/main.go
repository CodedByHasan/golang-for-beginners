package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "The quick brown fox jumped over the lazy dog."
	replaceString := "fox"

	result := strings.ReplaceAll(testString, replaceString, "hippo")

	fmt.Println(result)
}
