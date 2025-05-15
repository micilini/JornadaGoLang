package main

import (
    "fmt"
    "time"
)

var mensagemMostrada bool = false

func mostrarMensagem(texto string) {
    fmt.Println(texto)
    mensagemMostrada = true
}

func main() {
    go mostrarMensagem("Quando iremos aprender a usar goroutines, micilini?")

    // Loop para aguardar a execução da goroutine
    for !mensagemMostrada {
        time.Sleep(100 * time.Millisecond) // Pequeno delay para evitar alto consumo de CPU
    }
    
    fmt.Println("Final da Main()")
}