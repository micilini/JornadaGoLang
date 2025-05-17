package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	fmt.Println("Iniciando tarefa longa...")

	select {
	case <-time.After(5 * time.Second): // Simula tarefa que demora 5 segundos
		fmt.Println("✅ Tarefa concluída com sucesso.")
	case <-ctx.Done(): // Timeout ou cancelamento
		fmt.Println("⛔ Tarefa cancelada:", ctx.Err())
	}
}

func main() {
	// Cria um contexto com timeout de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Libera recursos quando terminar

	longRunningTask(ctx)
}