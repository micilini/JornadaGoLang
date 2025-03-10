package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Erro ao obter o caminho do executável:", err)
		return
	}

	fmt.Println("Caminho do executável:", exePath)
	fmt.Println("Diretório do executável:", filepath.Dir(exePath))
}