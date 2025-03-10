package main

import (
	"fmt"
	"os"
	"os/exec"
)

func listarArquivos() {
	cmd := exec.Command("cmd", "/C", "dir") // Windows
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao listar diretórios:", err)
	}
}

func main() {
	listarArquivos()
}