package main

import (
    "fmt"
    "sync"
)

func main() {
    var (
        counter int
        mu      sync.Mutex
        wg      sync.WaitGroup
    )

    numWorkers := 5
    wg.Add(numWorkers)

    for i := 1; i <= numWorkers; i++ {
        go func(id int) {
            defer wg.Done()

            mu.Lock()         // bloqueia o acesso ao counter
            defer mu.Unlock() // libera quando sair do bloco

            // seção crítica
            counter++
            fmt.Printf("Goroutine %d incrementou contador para %d\n", id, counter)
        }(i)
    }

    wg.Wait() // espera todas as goroutines terminarem
    fmt.Println("Valor final do contador:", counter)
}