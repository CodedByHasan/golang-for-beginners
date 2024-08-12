package main

import "fmt"

func factorial(n int) int {
	// base case
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var a int = 5
	fmt.Printf("%d!=%d\n", a, factorial(a))
}
