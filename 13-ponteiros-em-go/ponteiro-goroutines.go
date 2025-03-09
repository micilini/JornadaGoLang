package main

import (
    "fmt"
    "sync"
)

func incrementar(contador *int, wg *sync.WaitGroup) {
    defer wg.Done()
    *contador++
}

func main() {
    var wg sync.WaitGroup
    contador := 0

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go incrementar(&contador, &wg)
    }

    wg.Wait()
    fmt.Println("Valor final:", contador) // Saída: 5
}