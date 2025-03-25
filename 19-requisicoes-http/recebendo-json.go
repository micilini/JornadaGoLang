package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Estrutura do JSON esperado
type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Estrutura do JSON de resposta
type ResponseData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lendo o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Decodificando JSON recebido
	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Criando resposta JSON
	response := ResponseData{
		Message: fmt.Sprintf("Bem-vindo, %s! Seu email é %s.", requestData.Name, requestData.Email),
		Status:  200,
	}

	// Convertendo para JSON
	responseJSON, _ := json.Marshal(response)

	// Definindo cabeçalho da resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/api", handler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}