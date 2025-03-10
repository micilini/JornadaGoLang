package main

import (
	"fmt"
	"os"
)

func main() {
	// Cria a estrutura de diretórios recursivamente
	err := os.MkdirAll("minha-pasta/subpasta", 0755)
	if err != nil {
		fmt.Println("Erro ao criar subpasta:", err)
		return
	}
	fmt.Println("Subpasta criada em 'minha-pasta/subpasta'!")
}
