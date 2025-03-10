package main

import (
	"fmt"
	"os"
)

func removeEnvVar(key string) {
	err := os.Unsetenv(key)
	if err != nil {
		fmt.Println("Erro ao remover a variável:", err)
		return
	}
	fmt.Printf("Variável '%s' removida com sucesso!\n", key)
}

func main() {
	os.Setenv("TEMP_VAR", "Valor temporário")
	fmt.Println("Antes de remover:", os.Getenv("TEMP_VAR"))

	removeEnvVar("TEMP_VAR")

	fmt.Println("Após remover:", os.Getenv("TEMP_VAR"))
}