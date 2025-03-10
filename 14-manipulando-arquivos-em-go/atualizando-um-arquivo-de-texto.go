package main

import (
	"fmt"
	"os"
)

func main() {
	// Abrir em modo de apêndice (append) para adicionar sem sobrescrever
	file, err := os.OpenFile("meu-texto.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo para adicionar conteúdo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Essa é uma nova linha adicionada.\n")
	if err != nil {
		fmt.Println("Erro ao adicionar conteúdo:", err)
		return
	}

	fmt.Println("Conteúdo adicionado ao 'meu-texto.txt' com sucesso!")
}