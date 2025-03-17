package main

import (
	"encoding/json"
	"fmt"
)

func detectarJSON(jsonData []byte) {
	// Criamos uma variável do tipo `interface{}` para armazenar os dados
	var temp interface{}

	// Fazemos o parse do JSON
	err := json.Unmarshal(jsonData, &temp)
	if err != nil {
		fmt.Println("Erro ao fazer o parse do JSON:", err)
		return
	}

	// Verificamos o tipo da estrutura usando type assertion
	switch temp.(type) {
	case []interface{}:
		fmt.Println("O JSON é uma LISTA (array).")
	case map[string]interface{}:
		fmt.Println("O JSON é um OBJETO (map).")
	default:
		fmt.Println("Formato desconhecido.")
	}
}

func main() {
	// JSON de exemplo (lista)
	jsonLista := []byte(`[{"nome": "Carlos", "idade": 40}, {"nome": "Maria", "idade": 35}]`)

	// JSON de exemplo (objeto)
	jsonObjeto := []byte(`{"empresa": "TechCorp", "fundacao": 1999}`)

	fmt.Println("Detectando JSON Lista:")
	detectarJSON(jsonLista)

	fmt.Println("\nDetectando JSON Objeto:")
	detectarJSON(jsonObjeto)
}