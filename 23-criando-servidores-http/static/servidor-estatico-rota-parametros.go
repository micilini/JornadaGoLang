package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Função para tratar a rota dinâmica de um produto
func produtoHandler(w http.ResponseWriter, r *http.Request) {
	// Capturando o parâmetro 'id' da URL
	// Exemplo de URL: /produto/12345
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID de produto inválido", http.StatusBadRequest)
		return
	}

	idProduto := parts[2] // Pegando o ID do produto da URL

	// Aqui você pode validar se o idProduto existe no banco de dados, por exemplo
	// Para fins de exemplo, vamos apenas simular a resposta
	fmt.Fprintf(w, "Mostrando o produto com ID: %s", idProduto)
}

func main() {
	// Rota para mostrar informações de um produto específico
	http.HandleFunc("/produto/", produtoHandler)

	log.Println("Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}