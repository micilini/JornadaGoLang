package main

import (
	"fmt"
	"os"
)

func main() {
	// Alterar para o diretório "minha-pasta"
	err := os.Chdir("minha-pasta")
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return
	}

	// Confirmação do diretório alterado
	newDir, _ := os.Getwd()
	fmt.Println("Diretório atual após mudança:", newDir)
}