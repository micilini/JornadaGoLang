// internal/http/handlers/user_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "apis-em-go/internal/services"
)

// UserHandler agrupa métodos que respondem às rotas da API de usuários.
type UserHandler struct {
    Service *services.UserService
}

// NewUserHandler cria um UserHandler com o UserService injetado.
func NewUserHandler(svc *services.UserService) *UserHandler {
    return &UserHandler{Service: svc}
}

// respondJSON define cabeçalhos e serializa resposta JSON.
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(status)
    if payload != nil {
        json.NewEncoder(w).Encode(payload)
    }
}

// parseID extrai o ID do final da URL (por exemplo, /users/123).
func parseID(path string) (int, error) {
    // Path esperado: "/users/{id}"
    parts := strings.Split(path, "/")
    if len(parts) < 3 {
        return 0, http.ErrNotSupported
    }
    return strconv.Atoi(parts[2])
}

// ListUsersHandler lida com GET /users
func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    users, err := h.Service.ListUsers()
    if err != nil {
        respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    respondJSON(w, http.StatusOK, users)
}

// GetUserHandler lida com GET /users/{id}
func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    id, err := parseID(r.URL.Path)
    if err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
        return
    }

    user, err := h.Service.GetUser(id)
    if err != nil {
        respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
        return
    }
    respondJSON(w, http.StatusOK, user)
}

// CreateUserHandler lida com POST /users
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
        return
    }

    created, err := h.Service.CreateUser(input.Name, input.Email)
    if err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }
    respondJSON(w, http.StatusCreated, created)
}

// UpdateUserHandler lida com PUT /users/{id}
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    id, err := parseID(r.URL.Path)
    if err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
        return
    }

    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
        return
    }

    updated, err := h.Service.UpdateUser(id, input.Name, input.Email)
    if err != nil {
        if err.Error() == "usuário não encontrado" {
            respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
        } else {
            respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
        return
    }
    respondJSON(w, http.StatusOK, updated)
}

// DeleteUserHandler lida com DELETE /users/{id}
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    id, err := parseID(r.URL.Path)
    if err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
        return
    }

    if err := h.Service.DeleteUser(id); err != nil {
        respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
        return
    }
    respondJSON(w, http.StatusNoContent, nil)
}