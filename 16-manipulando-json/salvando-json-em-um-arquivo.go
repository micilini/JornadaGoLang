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
	// Criando uma struct
	pessoa := Pessoa{Nome: "Bruno", Idade: 30}

	// Criando o arquivo
	file, err := os.Create("pessoa.json")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Criando o encoder e escrevendo no arquivo
	encoder := json.NewEncoder(file)
	err = encoder.Encode(pessoa)
	if err != nil {
		fmt.Println("Erro ao escrever JSON:", err)
		return
	}

	fmt.Println("Arquivo JSON salvo com sucesso!")
}