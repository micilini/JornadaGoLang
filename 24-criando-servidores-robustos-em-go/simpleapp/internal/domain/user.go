// internal/domain/user.go
package domain

// User representa um usuário no sistema.
type User struct {
	Name  string
	Email string
}

// UserRepository define a interface para persistir/recuperar Users.
// É o “port” na Hexagonal/Clean Architecture.
type UserRepository interface {
	// GetAll retorna todos os usuários cadastrados.
	GetAll() ([]User, error)
	// Add persiste um novo usuário.
	Add(u User) error
}