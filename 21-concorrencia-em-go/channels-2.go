package main

import (
    "fmt"
    "time"
)

// Mostra um texto, dorme 5 segundos e envia um sinal para o canal indicando que terminou
func mostreAlgo(text string, done chan bool) {
    fmt.Println(text)
    time.Sleep(5 * time.Second) // pausa de 5 segundos
    done <- true                // sinal de conclusão
}

func main() {
    done := make(chan bool)    // canal de sinalização
    go mostreAlgo("Micilini Roll", done)
    <-done                     // bloqueia até receber o sinal de done
    fmt.Println("Saindo da função principal...")
}