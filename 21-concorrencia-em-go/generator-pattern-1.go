package main

import (
	"fmt"
)

// Função generator que retorna um canal com números de 1 a n
func generateNumbers(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch) // Fecha o canal após enviar os dados
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	for num := range generateNumbers(5) {
		fmt.Println(num)
	}
}