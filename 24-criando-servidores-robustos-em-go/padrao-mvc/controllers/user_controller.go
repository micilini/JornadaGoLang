// controllers/user_controller.go
package controllers

import (
	"html/template"
	"net/http"
	"padrao-mvc/models"
)

// UserController agrupa os métodos que manipulam usuários
type UserController struct {
	Templates *template.Template
}

// NewUserController recebe a instância de templates (pré-carregados) e retorna o controller
func NewUserController(tmpl *template.Template) *UserController {
	return &UserController{Templates: tmpl}
}

// ListUsers exibe a página principal ("/") com a lista de usuários cadastrados
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	// Recupera todos os usuários do Model
	allUsers := models.GetAll()

	// Define o cabeçalho de resposta como HTML UTF-8
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Renderiza o template "home.html", passando o slice de usuários como contexto
	if err := uc.Templates.ExecuteTemplate(w, "home.html", allUsers); err != nil {
		// Em caso de erro na renderização, retorna status 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddUser trata tanto GET quanto POST em "/add":
// - GET: exibe o formulário (add.html)
// - POST: processa o formulário, adiciona o usuário e redireciona para "/"
func (uc *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	// Se for GET, apenas renderiza o formulário para adicionar usuário
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := uc.Templates.ExecuteTemplate(w, "add.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Caso seja POST, processa os dados enviados pelo formulário
	if r.Method == http.MethodPost {
		// Faz o parse dos dados do form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Falha ao ler o formulário", http.StatusBadRequest)
			return
		}

		// Obtém os campos "name" e "email"
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Cria um novo User e adiciona ao Model
		newUser := models.User{Name: name, Email: email}
		models.Add(newUser)

		// Após adicionar, redireciona para a rota "/" (lista de usuários)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Se for outro método (PUT, DELETE etc.), retorna 405 Method Not Allowed
	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}