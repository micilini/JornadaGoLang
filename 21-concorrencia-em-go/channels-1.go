
package main

import (
	"fmt"
)

// It prints the text and sends a signal to the channel when done. Mostra um texto e envia um sinal para o canal indicando que a execução da função foi finalizada
func mostreAlgo(text string, done chan bool) {
	fmt.Println(text)
	done <- true // Envie um sinal (verdadeiro) no canal 'concluído' para indicar a conclusão
}

func main() {
	done := make(chan bool) // Cria um novo canal do tipo Bool chamado done

	go mostreAlgo("Micilini Roll", done)

	<-done // Bloqueia a execução das próximas lógicas abaixo até a função mostreAlgo dizer ao Channel que a execução foi finalizada

	fmt.Println("Saindo da função principal...")
}