package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go printNumbers(&wg)
    go printLetters(&wg)
    wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
    for i := 0; i <= 5; i++ {
        fmt.Println(i)
    }
    wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
    for i := 97; i <= 101; i++ {
        fmt.Println(string(rune(i)))
    }
    wg.Done()
}
