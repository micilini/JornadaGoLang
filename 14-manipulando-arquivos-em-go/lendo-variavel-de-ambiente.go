package main

import (
	"fmt"
	"os"
)

func getEnvVar(key string) {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("A variável de ambiente '%s' não está definida.\n", key)
	} else {
		fmt.Printf("Valor de '%s': %s\n", key, value)
	}
}

func main() {
	getEnvVar("USERNAME") // No Windows, geralmente retorna o nome do usuário
	getEnvVar("PATH")     // Exibe o PATH do sistema
}