package main

import (
	"fmt"
	"os"
)

func setEnvVar(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		fmt.Println("Erro ao definir a variável de ambiente:", err)
		return
	}
	fmt.Printf("Variável '%s' definida com sucesso!\n", key)
}

func main() {
	setEnvVar("MINHA_VARIAVEL", "12345")
	fmt.Println("MINHA_VARIAVEL:", os.Getenv("MINHA_VARIAVEL"))
}