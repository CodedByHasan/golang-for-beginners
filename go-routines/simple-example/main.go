// Concurrent program that calculates the square
// of a number (from 1 to 10 000) and displays it
// to the screen. This program takes less than 0.5
// seconds to execute!
package main

import (
    "fmt"
    "time"
)

func calculateSquare(i int) {
    time.Sleep(1 * time.Second)
    fmt.Println(i * i)
}

func main() {
    start := time.Now()

    for i := 1; i <= 10000; i++ {
        // goroutine
        go calculateSquare(i)
    }

    elapsed := time.Since(start)
    // need to Sleep for ~1 second to see output
    // of calculateSquare
    time.Sleep(1 * time.Second)
    fmt.Println("Function took: ", elapsed)
}
