// internal/domain/user.go
package domain

// User representa a entidade de usuário na API.
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// UserRepository define os métodos que qualquer fonte de dados deve implementar
// para armazenar e recuperar usuários.
type UserRepository interface {
    GetAll() ([]User, error)
    GetByID(id int) (*User, error)
    Create(u User) (*User, error)
    Update(id int, u User) (*User, error)
    Delete(id int) error
}