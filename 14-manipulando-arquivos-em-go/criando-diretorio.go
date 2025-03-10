package main

import (
	"fmt"
	"os"
)

func main() {
	// Cria o diretório "minha-pasta"
	err := os.Mkdir("minha-pasta", 0755)
	if err != nil {
		fmt.Println("Erro ao criar o diretório:", err)
		return
	}
	fmt.Println("Diretório 'minha-pasta' criado com sucesso!")
}
