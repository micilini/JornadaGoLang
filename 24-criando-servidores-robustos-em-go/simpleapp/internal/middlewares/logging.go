// internal/middlewares/logging.go
package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware grava no console o método, path e tempo de execução de cada requisição.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Chama o próximo handler na cadeia
		next.ServeHTTP(w, r)
		// Após a resposta, registra dados
		duration := time.Since(start)
		log.Printf("%s %s → %v\n", r.Method, r.URL.Path, duration)
	})
}