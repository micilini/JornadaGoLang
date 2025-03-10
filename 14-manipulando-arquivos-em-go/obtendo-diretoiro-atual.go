package main

import (
	"fmt"
	"os"
)

func main() {
	// Obter o diretório atual
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter diretório atual:", err)
		return
	}
	fmt.Println("Diretório atual:", currentDir)
}
