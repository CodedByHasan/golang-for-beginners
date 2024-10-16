package main

import (
	"fmt"
	"os"
)

var PATH string = "../test.txt"

func main() {
	fileInfo, err := os.Stat(PATH)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File name: %s\n", fileInfo.Name())
	fmt.Printf("File size: %d\n", fileInfo.Size())
	fmt.Printf("File permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Is it a directory? %t\n", fileInfo.IsDir())
}
