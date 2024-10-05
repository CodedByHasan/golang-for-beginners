package main

import (
    "fmt"
    "sync"
    "time"
)

// Numbers are printed in random order
// because go-routines are non-deterministic
func calculateSquare(i int, wg *sync.WaitGroup) {
    fmt.Println(i * i)
    wg.Done() // Decrement Wait Group counter by 1
}

func main() {
    var wg sync.WaitGroup
    start := time.Now()
    wg.Add(10)

    for i := 0; i < 10; i++ {
        go calculateSquare(i, &wg)
    }

    elapsed := time.Since(start)
    wg.Wait() // Block until Wait Group counter = 0
    fmt.Println("Function took ", elapsed)
}
