package main

import (
	"fmt"
)

// 3 funções que imprimem um texto e enviam um sinal de conclusão
func mostreAlgo(text string, done chan bool) {
	fmt.Println(text)
	done <- true
}

func mostreAlgo2(text string, done chan bool) {
	fmt.Println(text)
	done <- true
}

func mostreAlgo3(text string, done chan bool) {
	fmt.Println(text)
	done <- true
}

func main() {
	// canal bufferizado com capacidade para 2 sinais
	done := make(chan bool, 2)

	// Disparamos 3 goroutines
	go mostreAlgo("Chamou mostreAlgo", done)
	go mostreAlgo2("Chamou mostreAlgo2", done)
	go mostreAlgo3("Chamou mostreAlgo3", done)

	<-done
	<-done
	<-done

	fmt.Println("Saindo da função principal...")
}
