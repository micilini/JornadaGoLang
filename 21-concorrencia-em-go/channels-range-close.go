package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch) // Importante: fecha o canal para encerrar o range
    }()

    for val := range ch {
        fmt.Println("Recebido:", val)
    }

    fmt.Println("Canal fechado, fim do range.")
}