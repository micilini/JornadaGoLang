package main

import "fmt"

// Definindo uma struct chamada Pessoa
type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    // Criando uma instância da struct
    p := Pessoa{
        Nome:  "João",
        Idade: 30,
    }
    fmt.Println(p) // Saída: {João 30}
}
