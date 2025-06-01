// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"padrao-mvc/controllers"
)

func main() {
	// 1. Carrega todos os templates dentro de "templates/*.html"
	//    O template.Must faz panic se algum arquivo estiver com erro de sintaxe
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// 2. Cria o UserController, injetando a instância de templates
	userController := controllers.NewUserController(tmpl)

	// 3. Registra as rotas, associando cada caminho ao método do controller
	SetupRoutes(userController)

	// 4. Inicia o servidor na porta 3000
	log.Println("Servidor rodando em http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}