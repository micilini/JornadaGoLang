package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

	// Acessa as variáveis
	dbUser := os.Getenv("NOME")
	dbPass := os.Getenv("RANK")
	apiKey := os.Getenv("SITE_APP")

	fmt.Println("Usuário:", dbUser)
	fmt.Println("Rank:", dbPass)
	fmt.Println("URL:", apiKey)
}