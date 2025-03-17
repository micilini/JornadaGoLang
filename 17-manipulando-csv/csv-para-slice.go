package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Abrir o arquivo CSV
	file, err := os.Open("dados.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um leitor CSV
	reader := csv.NewReader(file)

	// Ler todo o conteúdo do arquivo CSV para um slice de slices
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler arquivo CSV:", err)
		return
	}

	// Exibir os dados lidos
	for _, record := range records {
		fmt.Println(record) // Cada "record" é um slice de strings
	}
}