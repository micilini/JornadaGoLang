package main

import (
    "fmt"
    "time"
)

func mostrarMensagem(texto string) {
    fmt.Println(texto)
}

func main() {
    go mostrarMensagem("Quando iremos aprender a usar goroutines, micilini?")

    // O comando abaixo meio que congela a execução do programa neste ponto, antes de executar os próximos comandos:
    time.Sleep(2 * time.Second)// Congela em 2 segundos, tempo suficiente para o goroutine mostrar uma mensagem no terminal
    
    fmt.Println("Final da Main()")
}