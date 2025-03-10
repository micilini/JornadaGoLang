package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erro ao obter o diretório home:", err)
		return
	}

	fmt.Println("Diretório Home do usuário:", homeDir)
}