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

    go buy(channel, wg)
    fmt.Println("Sent all data to the channel")
    close(channel)

    wg.Done()
}

func buy(channel chan int, wg *sync.WaitGroup) {
    fmt.Println("Waiting for data...")

    // Using a for-range to receive values from sell channel 
    for val := range channel {
        fmt.Printf("Received: %d\n", val)
    }

    wg.Done()
}
