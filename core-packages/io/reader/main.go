package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World")

	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]), err)

		if err != nil {
			fmt.Println("Error...")
			break
		}
	}
}
