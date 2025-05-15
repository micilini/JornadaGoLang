package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup, counter *int) {
    defer wg.Done()
    *counter++  // data race aqui!
    fmt.Printf("Goroutine %d: counter = %d\n", id, *counter)
}

func main() {
    var (
        counter int
        wg      sync.WaitGroup
    )

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg, &counter)
    }

    wg.Wait()
    fmt.Println("Valor final do contador:", counter)
}