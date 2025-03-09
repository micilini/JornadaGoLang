package main

import "fmt"

// Pessoa representa uma estrutura com nome, idade e cidade.
type Pessoa struct {
    Nome   string
    Idade  int
    Cidade string
}

// NovaPessoa recebe valores básicos e retorna uma struct Pessoa.
func NovaPessoa(nome string, idade int, cidade string) Pessoa {
    return Pessoa{
        Nome:   nome,
        Idade:  idade,
        Cidade: cidade,
    }
}

func main() {
    pessoa := NovaPessoa("Ana", 28, "Rio de Janeiro")
    fmt.Printf("Pessoa: %+v\n", pessoa)
}
