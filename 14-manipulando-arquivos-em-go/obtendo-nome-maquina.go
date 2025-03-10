package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Erro ao obter o nome da máquina:", err)
		return
	}

	fmt.Println("Nome da máquina:", hostname)
}