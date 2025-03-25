package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Criando um cliente HTTP
	client := &http.Client{}

	// Criando a requisição manualmente
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Adicionando headers personalizados
	req.Header.Set("Authorization", "Bearer MEU_TOKEN_AQUI")
	req.Header.Set("Custom-Header", "MeuValorPersonalizado")

	// Enviando a requisição
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