// routes.go
package main

import (
	"net/http"
	"padrao-mvc/controllers"
)

// SetupRoutes recebe a instância de UserController e associa cada rota ao handler correspondente
func SetupRoutes(uc *controllers.UserController) {
	// Rota principal "/" → lista de usuários
	http.HandleFunc("/", uc.ListUsers)

	// Rota "/add" → formulário e processamento de POST
	http.HandleFunc("/add", uc.AddUser)

	// (Opcional) servir arquivos estáticos caso você queira adicionar CSS, JS, imagens em /static/
//  fs := http.FileServer(http.Dir("public"))
//  http.Handle("/static/", http.StripPrefix("/static/", fs))
}