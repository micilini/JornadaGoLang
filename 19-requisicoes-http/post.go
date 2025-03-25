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

	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))
}