package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON desconhecido
	jsonData := `{"nome":"Carlos","idade":40,"profissao":"Engenheiro"}`

	// Criando um mapa para armazenar os dados
	var data map[string]interface{}

	// Convertendo JSON para mapa
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Exibindo os valores
	fmt.Println("Nome:", data["nome"])
	fmt.Println("Idade:", data["idade"])
	fmt.Println("Profissão:", data["profissao"])
}