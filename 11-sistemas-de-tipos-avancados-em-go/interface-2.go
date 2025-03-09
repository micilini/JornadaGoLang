package main

import "fmt"

// Definindo a interface Exibivel
type Exibivel interface {
    MostrarNome()
    MostrarIdade()
}

// Definindo a struct Pessoa
type Pessoa struct {
    Nome  string
    Idade int
}

// Implementando o método MostrarNome para a struct Pessoa
func (p Pessoa) MostrarNome() {
    fmt.Println("Nome:", p.Nome)
}

// Implementando o método MostrarIdade para a struct Pessoa
func (p Pessoa) MostrarIdade() {
    fmt.Println("Idade:", p.Idade)
}

func main() {
    // Criando uma instância de Pessoa
    pessoa := Pessoa{
        Nome:  "Carlos",
        Idade: 30,
    }

    // Criando uma variável do tipo Exibivel e atribuindo a pessoa
    var e Exibivel = pessoa

    // Chamando os métodos da interface
    e.MostrarNome()
    e.MostrarIdade()
}