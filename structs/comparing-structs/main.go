package main

import (
	"fmt"
)

type StructA struct {
	x int
}

func main() {
	s1 := StructA{x: 5}
	s2 := StructA{x: 6}
	s3 := StructA{x: 5}

	if s1 != s2 {
		fmt.Println("s1 and s2 have different values.")
	}
	if s1 == s3 {
		fmt.Println("s1 is the same as s3.")
	}

}
