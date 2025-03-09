package main

import "fmt"

// Definindo uma struct chamada Pessoa
type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
	p := Pessoa{"João", 28}

	// Acessando campos
	fmt.Println(p.Nome) // Saída: João

	// Modificando campos
	p.Idade = 29
	fmt.Println(p.Idade) // Saída: 29
}