package main

import (
	"fmt"
	"os"
)

var PATH string = "./core-packages/os/test2.txt"

func main() {
	file, err := os.OpenFile(PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
	}

	// delay closing of file until end of program
	defer file.Close()

	_, err = file.WriteString("Writing to file using WriteString!")

	if err != nil {
		fmt.Println(err)
	}
}
