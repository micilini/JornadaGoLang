// cmd/api/main.go
package main

import (
    "log"
    "net/http"

    "apis-em-go/internal/http/handlers"
    "apis-em-go/internal/http/middlewares"
    "apis-em-go/internal/repositories"
    "apis-em-go/internal/services"
)

func main() {
    // 1. Configurações iniciais (por ex., leitura de variáveis de ambiente) poderiam ser aqui

    // 2. Instancia o repositório (in-memory)
    userRepo := repositories.NewUserInMemory()

    // 3. Cria o serviço de usuário
    userSvc := services.NewUserService(userRepo)

    // 4. Cria o handler, injetando o UserService
    userHandler := handlers.NewUserHandler(userSvc)

    // 5. Configura mux (roteador) e registra rotas com middleware de logging
    mux := http.NewServeMux()

// "/users" => GET (listar) ou POST (criar)
mux.Handle("/users",
    middlewares.LoggingMiddleware(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case http.MethodGet:
                userHandler.ListUsersHandler(w, r)
            case http.MethodPost:
                userHandler.CreateUserHandler(w, r)
            default:
                http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
            }
        }),
    ),
)

// "/users/" => GET, PUT ou DELETE em "/users/{id}"
mux.Handle("/users/",
    middlewares.LoggingMiddleware(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case http.MethodGet:
                userHandler.GetUserHandler(w, r)
            case http.MethodPut:
                userHandler.UpdateUserHandler(w, r)
            case http.MethodDelete:
                userHandler.DeleteUserHandler(w, r)
            default:
                http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
            }
        }),
    ),
)

    // 6. Inicia o servidor na porta 3000
    addr := ":3000"
    log.Println("API rodando em http://localhost" + addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatal("Erro ao iniciar servidor:", err)
    }
}