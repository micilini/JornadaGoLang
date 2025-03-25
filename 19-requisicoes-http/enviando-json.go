package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Estrutura para o envio
type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Estrutura para a resposta
type ResponseData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	url := "http://localhost:8080/api"

	// Criando JSON para enviar
	requestData := RequestData{Name: "William", Email: "william@email.com"}
	jsonData, _ := json.Marshal(requestData)

	// Criando requisição HTTP
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Adicionando cabeçalhos
	req.Header.Set("Content-Type", "application/json")

	// Criando cliente HTTP e enviando requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lendo resposta
	body, _ := io.ReadAll(resp.Body)

	// Decodificando JSON de resposta
	var responseData ResponseData
	json.Unmarshal(body, &responseData)

	fmt.Println("Resposta do servidor:", responseData.Message)
}