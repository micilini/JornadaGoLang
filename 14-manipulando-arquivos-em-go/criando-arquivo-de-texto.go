package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("meu-texto.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Primeira linha do arquivo.\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo 'meu-texto.txt' criado com sucesso!")
}