package main

import (
    "fmt"
    "sync"
)

func main() {

    var wg sync.WaitGroup
    wg.Add(3)
    jobs := make(chan int)
    results := make(chan int)

    go func() {
        for j := range jobs {
            fmt.Println("received job", j)
            result := j * 2
            results <- result
        }
        close(results)
        wg.Done()
    }()

    go func() {
        for j := 1; j <= 3; j++ {
            jobs <- j
            fmt.Println("sent job", j)
        }
        close(jobs)
        wg.Done()
    }()

    go func() {
        for r := range results {
            fmt.Println("received result", r)
        }
        wg.Done()
    }()

    wg.Wait()
}
