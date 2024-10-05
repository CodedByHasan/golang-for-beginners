// This program spawns a goroutine closure in a loop
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(10)

    for i := 1; i <= 10; i++ {
        // This is a closure (function within a function)
        // pass in "i" as an argument so that it is well
        // defined at the point of calling the goroutine
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }

    wg.Wait()
    fmt.Println("Done")
}
