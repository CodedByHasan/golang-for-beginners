package main

import (
    "fmt"
    "time"
)

func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)

    go goChannelOne(channel1)
    go goChannelTwo(channel2)

    // output is non-deterministic because
    // select statement will select a
    // channel at random
    select {
    case val1 := <-channel1:
        fmt.Println(val1)
        break
    case val2 := <-channel2:
        fmt.Println(val2)
        break
    // Makes the select statement NON BLOCKING
    default:
        fmt.Println("Executed default block")
    }
    time.Sleep(1 * time.Second)
}

func goChannelOne(channel1 chan string) {
    channel1 <- "Channel 1"
}

func goChannelTwo(channel2 chan string) {
    channel2 <- "Channel 2"
}
