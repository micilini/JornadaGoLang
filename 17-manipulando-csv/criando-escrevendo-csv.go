package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("dados.csv")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escrever cabeçalho
	writer.Write([]string{"Nome", "Idade", "Profissão"})

	// Escrever dados
	writer.Write([]string{"Alice", "30", "Engenheira"})
	writer.Write([]string{"Bob", "25", "Desenvolvedor"})

	fmt.Println("Arquivo CSV criado com sucesso!")
}