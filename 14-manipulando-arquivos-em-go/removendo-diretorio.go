package main

import (
	"fmt"
	"os"
)

func main() {
	// Remove um diretório vazio
	err := os.Remove("minha-pasta/subpasta")
	if err != nil {
		fmt.Println("Erro ao remover diretório:", err)
		return
	}
	fmt.Println("Diretório 'subpasta' removido com sucesso!")
}