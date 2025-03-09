package main

import "fmt"

// Definindo a interface "Nomeado" com o método "MostrarNome"
type Nomeado interface {
    MostrarNome()
}

// Definindo a interface "Idade" com o método "MostrarIdade"
type Idade interface {
    MostrarIdade()
}

// Definindo a interface composta "Exibivel" que inclui "Nomeado" e "Idade"
type Exibivel interface {
    Nomeado  // Incluindo a interface Nomeado
    Idade    // Incluindo a interface Idade
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

    // Criando uma variável do tipo Exibivel, que é composta por Nomeado e Idade
    var e Exibivel = pessoa

    // Chamando os métodos da interface Exibivel
    e.MostrarNome()
    e.MostrarIdade()
}