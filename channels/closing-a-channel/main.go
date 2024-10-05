package main

import "fmt"

func main() {
    channel := make(chan int, 10)

    // Sending some values to the channel
    channel <- 1
    channel <- 2

    val, ok := <-channel
    fmt.Println(val, ok)

    // closing channel
    close(channel)

    // Closing an already closed channel will
    // cause a panic situation!
    // close(channel)

    // Sending data to a closed channel will cause
    // a panic situation!
    // channel <- 3

    val, ok = <-channel
    fmt.Println(val, ok)

    // Will return "0 false" because
    // channel is closed and there are
    // no more values to receive.
    val, ok = <-channel
    fmt.Println(val, ok)
}
