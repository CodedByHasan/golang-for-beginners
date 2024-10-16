package main

import (
	"fmt"
	"os"
)

var PATH string = "../test.txt"

func main() {
	file, err := os.Open(PATH)

	if err != nil {
		fmt.Println(err)
	}

	b := make([]byte, 16)

	for {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(b[0:n])
		fmt.Println(string(b[0:n]))
	}
}
