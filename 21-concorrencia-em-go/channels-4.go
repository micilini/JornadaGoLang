package main

import (
    "fmt"
)

// sendString envia a mensagem pelo canal e imprime um log
func sendString(ch chan string, msg string) {
    fmt.Println("Enviando:", msg)
    ch <- msg
}

func main() {
    // Cria um canal de strings
    msgChan := make(chan string)

    // Chama a função em uma goroutine separada
    go sendString(msgChan, "Olá, canal de strings!")

    // Recebe a mensagem enviada e armazena em 'received'
    received := <-msgChan
    fmt.Println("Recebido do canal:", received)
}