package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	// Abrindo o arquivo
	file, err := os.Open("pessoa.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criando uma variável para armazenar os dados
	var pessoa Pessoa

	// Criando um decoder e lendo os dados do arquivo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pessoa)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}

	// Exibindo os dados
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
}