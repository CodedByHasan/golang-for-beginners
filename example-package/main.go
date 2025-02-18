package main

// importing third party package
import (
	"fmt"
	"github.com/CodedByHasan/encryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("Hello World")
	fmt.Println(encryptedStr)
}
