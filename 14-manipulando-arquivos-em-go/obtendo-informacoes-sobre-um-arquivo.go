package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("meu-texto-renomeado.txt")
	if err != nil {
		fmt.Println("Erro ao obter informações do arquivo:", err)
		return
	}

	fmt.Println("Informações do arquivo:")
	fmt.Println("Nome:", info.Name())
	fmt.Println("Tamanho (bytes):", info.Size())
	fmt.Println("Permissões:", info.Mode())
	fmt.Println("Última modificação:", info.ModTime())
}