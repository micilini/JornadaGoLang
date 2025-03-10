package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func aguardarSinal() {
	canal := make(chan os.Signal, 1)
	signal.Notify(canal, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Aguardando CTRL+C ou SIGTERM...")
	sinalRecebido := <-canal
	fmt.Println("Sinal recebido:", sinalRecebido)
}

func main() {
	aguardarSinal()
}