// models/user.go
package models

// User representa um usuário cadastrado no sistema
type User struct {
	Name  string
	Email string
}

// slice que armazenará todos os usuários (em memória)
var users []User

// init inicializa a lista de usuários (vazia) ao carregar o pacote
func init() {
	users = []User{}
}

// GetAll retorna todos os usuários cadastrados
func GetAll() []User {
	return users
}

// Add adiciona um novo usuário à lista
func Add(u User) {
	users = append(users, u)
}