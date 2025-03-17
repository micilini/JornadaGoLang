package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON com uma lista de objetos
	jsonData := `[{"nome":"Carlos","idade":40,"profissao":"Engenheiro"},
	              {"nome":"Maria","idade":35,"profissao":"Médica"},
	              {"nome":"João","idade":28}]`

	// Criando um slice de maps
	var pessoas []map[string]interface{}

	// Convertendo JSON para slice de mapas
	err := json.Unmarshal([]byte(jsonData), &pessoas)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Iterando sobre os objetos do JSON
	for i, pessoa := range pessoas {
		fmt.Println("Pessoa", i+1)
		fmt.Println("Nome:", pessoa["nome"])
		fmt.Println("Idade:", pessoa["idade"])

		// Profissão pode não existir em todos os objetos
		if profissao, existe := pessoa["profissao"]; existe {
			fmt.Println("Profissão:", profissao)
		} else {
			fmt.Println("Profissão: Não informada")
		}
		fmt.Println("----------------------")
	}
}