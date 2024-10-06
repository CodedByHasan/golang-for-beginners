package main

import (
    "fmt"
    "time"
)

func main() {
    // anonymous method
    go func() {
        fmt.Println("In anonymous method")
    }()
    time.Sleep(1 * time.Second)
}
