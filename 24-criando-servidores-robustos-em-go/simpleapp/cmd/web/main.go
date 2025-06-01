// cmd/web/main.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"simpleapp/internal/http/handlers"
	"simpleapp/internal/middlewares"
	"simpleapp/internal/repositories"
	"simpleapp/internal/services"
)

func main() {
	// 1. Carrega todos os templates (*.html) no diretório "templates"
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// 2. Instancia o repositório in‐memory e o serviço
	userRepo := repositories.NewUserInMemory()
	userSvc := services.NewUserService(userRepo)

	// 3. Cria o handler, injetando o serviço e templates
	userHandler := handlers.NewUserHandler(userSvc, tmpl)

	// 4. Configura o mux e registra rotas com middleware de logging
	mux := http.NewServeMux()

	// Middleware de logging: LoggingMiddleware( handler )
	mux.Handle("/", middlewares.LoggingMiddleware(http.HandlerFunc(userHandler.List)))
	mux.Handle("/add", middlewares.LoggingMiddleware(http.HandlerFunc(userHandler.Add)))

	// 5. Inicia o servidor
	addr := ":3000"
	log.Println("Servidor rodando em http://localhost" + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}