package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func produtoHandler(w http.ResponseWriter, r *http.Request) {
	// Usando uma expressão regular para capturar o ID do produto
	// Exemplo de URL: /produto/12345 ou /produto/prod-12345
	re := regexp.MustCompile(`^/produto/([a-zA-Z0-9\-]+)$`)

	// Tentando combinar a URL com a expressão regular
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	// O primeiro item de matches é a string inteira, o segundo é o grupo capturado (ID do produto)
	idProduto := matches[1]

	// Lógica para buscar o produto pelo ID
	fmt.Fprintf(w, "Mostrando o produto com ID: %s", idProduto)
}

func main() {
	http.HandleFunc("/produto/", produtoHandler)

	log.Println("Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}