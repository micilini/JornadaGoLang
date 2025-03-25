package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	jsonData := `{"id": 1, "title": "Atualizado", "body": "Conteúdo atualizado", "userId": 1}`
	reqBody := bytes.NewBuffer([]byte(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, reqBody)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))
}