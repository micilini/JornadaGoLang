package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	jsonData := `{"title": "GoLang", "body": "Aprendendo Go!", "userId": 1}`
	reqBody := bytes.NewBuffer([]byte(jsonData))

	// Criando cliente HTTP
	client := &http.Client{}

	// Criando a requisição manualmente
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Adicionando headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer MEU_TOKEN_AQUI")
	req.Header.Set("Custom-Header", "MeuValorPersonalizado")

	// Enviando requisição
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lendo resposta
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))
}