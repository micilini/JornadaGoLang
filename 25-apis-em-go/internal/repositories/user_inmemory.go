// internal/repositories/user_inmemory.go
package repositories

import (
    "errors"
    "sync"

    "apis-em-go/internal/domain"
)

type userInMemory struct {
    mu      sync.RWMutex
    storage []domain.User
    nextID  int
}

// NewUserInMemory retorna uma instância de UserRepository que armazena dados em memória.
func NewUserInMemory() domain.UserRepository {
    return &userInMemory{
        storage: make([]domain.User, 0),
        nextID:  1,
    }
}

func (r *userInMemory) GetAll() ([]domain.User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    result := make([]domain.User, len(r.storage))
    copy(result, r.storage)
    return result, nil
}

func (r *userInMemory) GetByID(id int) (*domain.User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    for _, u := range r.storage {
        if u.ID == id {
            // Retorna uma cópia para evitar vazamento do slice interno
            usr := u
            return &usr, nil
        }
    }
    return nil, errors.New("usuário não encontrado")
}

func (r *userInMemory) Create(u domain.User) (*domain.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    u.ID = r.nextID
    r.nextID++
    r.storage = append(r.storage, u)

    created := u
    return &created, nil
}

func (r *userInMemory) Update(id int, u domain.User) (*domain.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    for i, existing := range r.storage {
        if existing.ID == id {
            u.ID = id
            r.storage[i] = u
            updated := u
            return &updated, nil
        }
    }
    return nil, errors.New("usuário não encontrado")
}

func (r *userInMemory) Delete(id int) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    for i, u := range r.storage {
        if u.ID == id {
            // Remove o elemento do slice
            r.storage = append(r.storage[:i], r.storage[i+1:]...)
            return nil
        }
    }
    return errors.New("usuário não encontrado")
}