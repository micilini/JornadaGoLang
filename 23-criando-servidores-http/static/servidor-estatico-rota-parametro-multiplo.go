package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Função para tratar a rota dinâmica de produto e SKU
func produtoHandler(w http.ResponseWriter, r *http.Request) {
	// Capturando os parâmetros 'id' e 'sku' da URL
	// Exemplo de URL: /produto/12345/sku1234
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Parâmetros inválidos", http.StatusBadRequest)
		return
	}

	idProduto := parts[2] // ID do produto
	sku := parts[3]       // SKU do produto

	// Aqui você pode validar se o idProduto e sku existem no banco de dados, por exemplo
	// Para fins de exemplo, vamos apenas simular a resposta
	fmt.Fprintf(w, "Mostrando o produto com ID: %s e SKU: %s", idProduto, sku)
}

func main() {
	// Rota para mostrar informações de um produto com ID e SKU
	http.HandleFunc("/produto/", produtoHandler)

	log.Println("Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}