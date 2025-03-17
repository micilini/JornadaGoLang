package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Definir a struct para representar as pessoas
type Pessoa struct {
	ID       int
	Nome     string
	Idade    int
	Profissao string
}

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

	// Criar um slice para armazenar os dados convertidos em structs
	var pessoas []Pessoa

	// Ignorar o cabeçalho e ler as linhas
	for i, record := range records {
		if i == 0 { // Ignorar o cabeçalho
			continue
		}
		id, _ := strconv.Atoi(record[0])  // Converter string para int
		idade, _ := strconv.Atoi(record[2]) // Converter string para int

		// Criar e adicionar a struct
		pessoa := Pessoa{
			ID:       id,
			Nome:     record[1],
			Idade:    idade,
			Profissao: record[3],
		}
		pessoas = append(pessoas, pessoa)
	}

	// Exibir as pessoas
	for _, p := range pessoas {
		fmt.Printf("ID: %d, Nome: %s, Idade: %d, Profissão: %s\n", p.ID, p.Nome, p.Idade, p.Profissao)
	}
}