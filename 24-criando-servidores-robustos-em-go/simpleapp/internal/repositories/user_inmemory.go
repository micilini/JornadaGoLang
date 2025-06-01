// internal/repositories/user_inmemory.go
package repositories

import (
	"fmt"
	"sync"

	"simpleapp/internal/domain"
)

// userInMemory é a struct que implementa domain.UserRepository usando um slice em memória.
// Protegemos o slice com um mutex para concorrência.
type userInMemory struct {
	mu    sync.RWMutex
	store []domain.User
}

// NewUserInMemory cria uma instância de UserRepository em memória.
func NewUserInMemory() domain.UserRepository {
	return &userInMemory{
		store: make([]domain.User, 0),
	}
}

// GetAll retorna todos os usuários. Bloqueia mutex para leitura.
func (r *userInMemory) GetAll() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Retorna uma cópia do slice para evitar que chamador manipule diretamente o slice interno
	copySlice := make([]domain.User, len(r.store))
	copy(copySlice, r.store)
	return copySlice, nil
}

// Add adiciona um novo usuário. Bloqueia mutex para escrita.
func (r *userInMemory) Add(u domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Poderíamos, por exemplo, checar se já existe um e-mail igual:
	for _, existing := range r.store {
		if existing.Email == u.Email {
			return fmt.Errorf("usuário com e-mail %s já existe", u.Email)
		}
	}

	r.store = append(r.store, u)
	return nil
}