package main

import (
    "fmt"
    "time"
)

func send(id int, ch chan int) {
    fmt.Printf("Goroutine %d: antes de enviar\n", id)
    ch <- id
    fmt.Printf("Goroutine %d: depois de enviar\n", id)
}

func main() {
    ch := make(chan int, 2) // buffer de capacidade 2

    // Dispara 3 goroutines que tentam enviar
    go send(1, ch)
    go send(2, ch)
    go send(3, ch)

    // Dá um tempinho para vermos as duas primeiras passarem
    time.Sleep(500 * time.Millisecond)

    // Vamos ler apenas UMA vez
    fmt.Println("Main: lendo um valor ->", <-ch)

    // Aguarda mais um pouco para destravar a 3ª goroutine
    time.Sleep(500 * time.Millisecond)

    // Leitura final para esvaziar o canal
    fmt.Println("Main: lendo os restantes ->", <-ch, <-ch)
}