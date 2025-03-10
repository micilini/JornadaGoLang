package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("meu-texto.txt", "meu-texto-renomeado.txt")
	if err != nil {
		fmt.Println("Erro ao renomear o arquivo:", err)
		return
	}

	fmt.Println("Arquivo renomeado para 'meu-texto-renomeado.txt'.")
}