package main

import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

func (p Pessoa) Saudacao() string {
    return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos.", p.Nome, p.Idade)
}

func main() {
    pessoa := Pessoa{"Carlos", 30}
    fmt.Println(pessoa.Saudacao())
}