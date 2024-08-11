package main

import "fmt"

func main() {
	var i *int      // declare int pointer
	var str *string //declare string pointer

	// Printing value of pointers to stdout
	// will be <nill> - zero value of pointer
	fmt.Println(i)
	fmt.Println(str)
}
