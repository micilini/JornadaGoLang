package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func detectarJSONComDecoder(jsonData string) {
	// Criamos um decoder para processar o JSON em stream
	decoder := json.NewDecoder(strings.NewReader(jsonData))

	// Pegamos o primeiro token do JSON
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}

	// Verificamos se o primeiro token é um '[' (lista) ou '{' (objeto)
	switch token {
	case json.Delim('['):
		fmt.Println("O JSON é uma LISTA (array).")
	case json.Delim('{'):
		fmt.Println("O JSON é um OBJETO (map).")
	default:
		fmt.Println("Formato desconhecido.")
	}
}

func main() {
	// JSON de exemplo (lista)
	jsonLista := `[{"nome": "Carlos", "idade": 40}, {"nome": "Maria", "idade": 35}]`

	// JSON de exemplo (objeto)
	jsonObjeto := `{"empresa": "TechCorp", "fundacao": 1999}`

	fmt.Println("Detectando JSON Lista:")
	detectarJSONComDecoder(jsonLista)

	fmt.Println("\nDetectando JSON Objeto:")
	detectarJSONComDecoder(jsonObjeto)
}