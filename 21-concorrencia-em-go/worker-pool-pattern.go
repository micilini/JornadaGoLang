package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker: executa uma tarefa recebida via canal
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: iniciando trabalho %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // simula trabalho
		fmt.Printf("Worker %d: finalizou trabalho %d\n", id, job)
		results <- job * 2 // exemplo: resultado fictício
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const totalJobs = 10
	const numWorkers = 3

	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	// Inicia os workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Envia os jobs para os workers
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}
	close(jobs) // importante: sinaliza que não virão mais tarefas

	// Recebe os resultados
	for r := 1; r <= totalJobs; r++ {
		fmt.Printf("Resultado recebido: %d\n", <-results)
	}
}