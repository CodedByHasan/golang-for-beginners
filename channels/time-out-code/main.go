package main

import (
    "fmt"
    "time"
)

func main() {
    channel1 := make(chan int)
    go sendValue(channel1)

    select {
    case msg := <-channel1:
        fmt.Println(msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Select timeout")
    }
}

func sendValue(channel chan int) {
    time.Sleep(2 * time.Second)
    channel <- 10
}
