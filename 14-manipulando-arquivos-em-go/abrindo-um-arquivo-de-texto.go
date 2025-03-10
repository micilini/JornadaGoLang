package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("meu-texto.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("Conteúdo do arquivo:")
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler o conteúdo:", err)
		return
	}
	fmt.Println(string(content))
}