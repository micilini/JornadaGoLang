package main

import (
	"fmt"
	"os"
)

func validarEntrada(valor int) {
	if valor < 0 {
		fmt.Println("Erro: Valor negativo não permitido.")
		os.Exit(1) // Sai com código de erro 1
	}
	fmt.Println("Valor válido:", valor)
}

func main() {
	validarEntrada(-10)
	fmt.Println("Esse código não será executado.")
}