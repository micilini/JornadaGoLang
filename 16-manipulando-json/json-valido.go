package main

import (
	"encoding/json"
	"fmt"
)

func verificarJSON(jsonData []byte) {
	if json.Valid(jsonData) {
		fmt.Println("✅ O JSON é válido!")
	} else {
		fmt.Println("❌ O JSON é inválido!")
	}
}

func main() {
	// JSON válido (Objeto)
	jsonValido := []byte(`{"nome": "Carlos", "idade": 30}`)

	// JSON válido (Lista)
	jsonListaValida := []byte(`[{"nome": "Ana"}, {"nome": "Pedro"}]`)

	// JSON inválido (chave sem aspas e falta de vírgula)
	jsonInvalido := []byte(`{nome: "Carlos" "idade": 30}`)

	// JSON inválido (falta de fechamento do colchete)
	jsonListaInvalida := []byte(`[{"nome": "Ana", "idade": 25}`)

	fmt.Println("Verificando JSON Objeto:")
	verificarJSON(jsonValido)

	fmt.Println("\nVerificando JSON Lista:")
	verificarJSON(jsonListaValida)

	fmt.Println("\nVerificando JSON Inválido (Objeto):")
	verificarJSON(jsonInvalido)

	fmt.Println("\nVerificando JSON Inválido (Lista):")
	verificarJSON(jsonListaInvalida)
}