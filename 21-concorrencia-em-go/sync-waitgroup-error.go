package main

import (
    "fmt"
    "sync"
)

func task(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Task %d concluída\n", id)
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2) // Só adicionamos duas “esperas” ao WaitGroup

    go task(1, &wg)
    go task(2, &wg)
    go task(3, &wg) // Terceira goroutine

    wg.Wait()
    fmt.Println("Todas as tarefas esperadas concluídas")
}