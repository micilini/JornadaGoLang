package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Métodos permitidos:", resp.Header.Get("Allow"))
}