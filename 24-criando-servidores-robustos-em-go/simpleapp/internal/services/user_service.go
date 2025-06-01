// internal/services/user_service.go
package services

import (
	"errors"
	"strings"

	"simpleapp/internal/domain"
)

// UserService orquestra o uso de domain.UserRepository e aplica validações.
type UserService struct {
	Repo domain.UserRepository
}

// NewUserService cria uma instância de UserService com o repositório injetado.
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// ListUsers retorna todos os usuários.
// Propaga qualquer erro de Repo.GetAll.
func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.Repo.GetAll()
}

// CreateUser valida os dados e, se estiverem corretos, adiciona ao repositório.
// Retorna erro caso o e-mail esteja vazio, formato inválido, ou se Repo.Add falhar.
func (s *UserService) CreateUser(name, email string) error {
	// Validações básicas:
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return errors.New("nome é obrigatório")
	}
	if email == "" {
		return errors.New("email é obrigatório")
	}
	// Poderíamos adicionar validação de formato de e-mail aqui (regex), mas mantemos simples.

	// Cria entidade de domínio e persiste
	u := domain.User{Name: name, Email: email}
	return s.Repo.Add(u)
}