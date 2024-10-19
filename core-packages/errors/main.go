package main

import (
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("Only odd numbers allowed, got: %d", i)
	}
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation successful!")
}

func main() {
	err := process(3)
	checkError(err)

	err = process(2)
	checkError(err)
}
