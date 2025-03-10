package main

import (
	"fmt"
	"os"
)

func main() {
	// Permissões: 0664 -> Usuário e grupo com leitura/escrita, outros apenas leitura
	file, err := os.OpenFile("arquivo-custom.txt", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Escrevendo com permissões customizadas!\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo 'arquivo-custom.txt' criado com sucesso!")
}