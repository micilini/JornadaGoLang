package main

import (
	"log"
	"net/http"
)

func main() {
	// Rota para a página inicial
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Rota para a página 'sobre'
	http.HandleFunc("/sobre", func(w http.ResponseWriter, r *http.Request) {
		// Aqui você pode criar uma resposta personalizada, como enviar um HTML diferente
		http.ServeFile(w, r, "public/sobre.html")
	})

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
