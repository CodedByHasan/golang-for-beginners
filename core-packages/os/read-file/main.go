package main

import (
	"fmt"
	"os"
)

var PATH string = "../test.txt"

func main() {
	data, err := os.ReadFile(PATH)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))
}
