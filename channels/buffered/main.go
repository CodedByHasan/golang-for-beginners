package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    channel := make(chan int, 3)
    go sell(channel, &wg)
    wg.Wait()
}

func sell(channel chan int, wg *sync.WaitGroup) {
    channel <- 1
    channel <- 2
    channel <- 3
    // This will cause a deadlock because buffer size > 3
    // channel <- 4
    go buy(channel, wg)
    fmt.Println("Sent all data to the channel")

    wg.Done()
}

func buy(channel chan int, wg *sync.WaitGroup) {
    fmt.Println("Waiting for data...")

    fmt.Println("Received data: ", <-channel)
    fmt.Println("Received data: ", <-channel)
    fmt.Println("Received data: ", <-channel)

    // This will cause a deadlock because buffer size > 3
    // fmt.Println("Received data: ", <-channel)

    wg.Done()
}
