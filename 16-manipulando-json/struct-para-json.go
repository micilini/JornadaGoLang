package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	// Criando uma instância da struct
	pessoa := Pessoa{Nome: "Alice", Idade: 25}

	// Convertendo a struct para JSON
	jsonBytes, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	// Convertendo bytes para string e imprimindo
	fmt.Println(string(jsonBytes))
}