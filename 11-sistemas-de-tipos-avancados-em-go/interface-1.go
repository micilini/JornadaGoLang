package main

import "fmt"

// Definindo a struct Pessoa
type Pessoa struct {
    Nome  string
    Idade int
}

// Função para mostrar o nome da pessoa
func MostrarNome(p Pessoa) {
    fmt.Println("Nome:", p.Nome)
}

// Função para mostrar a idade da pessoa
func MostrarIdade(p Pessoa) {
    fmt.Println("Idade:", p.Idade)
}

func main() {
    // Criando uma instância de Pessoa
    pessoa := Pessoa{
        Nome:  "Carlos",
        Idade: 30,
    }

    // Chamando as funções para mostrar o nome e a idade
    MostrarNome(pessoa)
    MostrarIdade(pessoa)
}