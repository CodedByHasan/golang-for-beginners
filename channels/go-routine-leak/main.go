package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go leak(&wg)
    wg.Wait()
}

func leak(wg *sync.WaitGroup) {
    channel := make(chan int)

    // anonymous go routine will wait indefinitely
    // which will cause a go routing leak (deadlock)!
    go func() {
        val := <-channel
        fmt.Println(val)
        wg.Done()
    }()

    fmt.Println("Exiting leak method")
    wg.Done()
}
