package main

import (
	"fmt"
	"sync"
)

func mostreAlgo(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Sinaliza que esta Goroutine está pronta quando a função sai
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // Cria uma WaitGroup

	wg.Add(1) // Incrementa o contador WaitGroup
	go mostreAlgo("Micilini Roll", &wg)
	wg.Wait() // Espera que todas as GoRoutines terminem

	fmt.Println("Saindo da função principal...")
}