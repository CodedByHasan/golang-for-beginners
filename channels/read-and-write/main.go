// The channels in this program are Unbuffered Channels
// This type of channel needs a receiver as SOON as a
// message is emitted to the channel.
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    channel := make(chan string)

    go sell(channel, &wg)
    go buy(channel, &wg)

    wg.Wait()
}

// sends data to the channel
func sell(channel chan string, wg *sync.WaitGroup) {
    // sell() will BLOCK go routine until another channel receives data
    channel <- "SOME DATA"
    fmt.Println("Sent data to the channel")

    wg.Done()
}

// receive data from the channel
func buy(channel chan string, wg *sync.WaitGroup) {
    fmt.Println("Waiting for data...")

    // buy() will UNBLOCK sell() go routine
    val := <-channel
    fmt.Printf("Received data - %s\n", val)

    wg.Done()
}
