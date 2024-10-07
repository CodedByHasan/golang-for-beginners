package main

import (
    "fmt"
    "github.com/codedByHasan/golang-for-beginners/decrypt"
    "github.com/codedByHasan/golang-for-beginners/encrypt"
)

func main() {
    encryptStr := encrypt.Nimbus("Hello World")
    fmt.Println("Encrypted String: ", encryptStr)
    fmt.Println("Decrypting the string...")
    fmt.Println(decrypt.Nimbus(encryptStr))
}
