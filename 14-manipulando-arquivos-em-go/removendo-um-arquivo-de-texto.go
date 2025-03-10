package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("meu-texto-renomeado.txt")
	if err != nil {
		fmt.Println("Erro ao remover o arquivo:", err)
		return
	}

	fmt.Println("Arquivo 'meu-texto-renomeado.txt' removido com sucesso!")
}
