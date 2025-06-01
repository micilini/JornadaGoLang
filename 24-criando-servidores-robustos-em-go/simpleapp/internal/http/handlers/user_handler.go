// internal/http/handlers/user_handler.go
package handlers

import (
	"html/template"
	"net/http"

	"simpleapp/internal/services"
)

// UserHandler contém dependências para atender requisições relacionadas a User.
type UserHandler struct {
	Service   *services.UserService
	Templates *template.Template
}

// NewUserHandler recebe a instância de service e templates já carregados.
func NewUserHandler(svc *services.UserService, tmpl *template.Template) *UserHandler {
	return &UserHandler{Service: svc, Templates: tmpl}
}

// List exibe a página com lista de usuários (GET "/").
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// Chama o service para obter todos os usuários
	users, err := h.Service.ListUsers()
	if err != nil {
		http.Error(w, "Erro ao obter usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Define cabeçalho
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Renderiza o template home.html, passando o slice de domain.User como contexto
	if err := h.Templates.ExecuteTemplate(w, "home.html", users); err != nil {
		http.Error(w, "Erro ao renderizar template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Add exibe o formulário (GET "/add") e processa submissão (POST "/add").
func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Renderiza o formulário
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := h.Templates.ExecuteTemplate(w, "add.html", nil); err != nil {
			http.Error(w, "Erro ao renderizar template: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		// Processa os dados do form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Falha ao ler formulário: "+err.Error(), http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Chama o serviço para criar o usuário
		if err := h.Service.CreateUser(name, email); err != nil {
			// Em caso de erro (nome vazio, e‐mail vazio ou e‐mail duplicado), retornamos ao form com mensagem
			http.Error(w, "Falha ao adicionar usuário: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Sucesso → redireciona para List ("/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Se for outro método, devolve 405
	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}