package main

import (
	"fmt"
	"log"
	
	"os"
)

var PATH string = "./app.log"

func main() {

	file, err := os.OpenFile(PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(file)
	log.Println("Why so serious")
}
