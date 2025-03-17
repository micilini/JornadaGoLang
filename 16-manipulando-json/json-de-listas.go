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
	// JSON de uma lista de pessoas
	jsonData := `[{"nome":"Ana","idade":20},{"nome":"Pedro","idade":35}]`

	// Criando um slice para armazenar os dados
	var pessoas []Pessoa

	// Convertendo JSON para slice de structs
	err := json.Unmarshal([]byte(jsonData), &pessoas)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Exibindo os dados
	for _, pessoa := range pessoas {
		fmt.Println("Nome:", pessoa.Nome, "| Idade:", pessoa.Idade)
	}
}