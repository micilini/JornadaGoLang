// internal/http/middlewares/logging.go
package middlewares

import (
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware imprime método, caminho e duração da requisição.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)
        log.Printf("%s %s → %v\n", r.Method, r.URL.Path, duration)
    })
}