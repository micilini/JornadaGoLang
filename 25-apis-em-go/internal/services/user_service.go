// internal/services/user_service.go
package services

import (
    "errors"
    "regexp"
    "strings"

    "apis-em-go/internal/domain"
)

// UserService é responsável por regras de negócio e usa um UserRepository.
type UserService struct {
    Repo domain.UserRepository
}

// NewUserService cria um UserService com o repositório injetado.
func NewUserService(repo domain.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

// ListUsers retorna todos os usuários.
func (s *UserService) ListUsers() ([]domain.User, error) {
    return s.Repo.GetAll()
}

// GetUser retorna um usuário pelo ID.
func (s *UserService) GetUser(id int) (*domain.User, error) {
    return s.Repo.GetByID(id)
}

// CreateUser valida e então cria um novo usuário.
func (s *UserService) CreateUser(name, email string) (*domain.User, error) {
    name = strings.TrimSpace(name)
    email = strings.TrimSpace(email)

    if name == "" {
        return nil, errors.New("nome é obrigatório")
    }
    if email == "" {
        return nil, errors.New("email é obrigatório")
    }
    // Validação simples de e-mail
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !re.MatchString(email) {
        return nil, errors.New("formato de email inválido")
    }

    user := domain.User{
        Name:  name,
        Email: email,
    }
    return s.Repo.Create(user)
}

// UpdateUser atualiza um usuário existente.
func (s *UserService) UpdateUser(id int, name, email string) (*domain.User, error) {
    name = strings.TrimSpace(name)
    email = strings.TrimSpace(email)

    if name == "" {
        return nil, errors.New("nome é obrigatório")
    }
    if email == "" {
        return nil, errors.New("email é obrigatório")
    }
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !re.MatchString(email) {
        return nil, errors.New("formato de email inválido")
    }

    user := domain.User{
        Name:  name,
        Email: email,
    }
    return s.Repo.Update(id, user)
}

// DeleteUser remove um usuário pelo ID.
func (s *UserService) DeleteUser(id int) error {
    return s.Repo.Delete(id)
}