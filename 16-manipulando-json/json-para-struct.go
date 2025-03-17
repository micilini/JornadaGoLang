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
	// JSON de entrada (simulação de um arquivo ou API)
	jsonData := `{"nome":"Alice","idade":25}`

	// Criando uma variável para armazenar os dados
	var pessoa Pessoa

	// Convertendo JSON para struct
	err := json.Unmarshal([]byte(jsonData), &pessoa)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Exibindo os dados
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
}
