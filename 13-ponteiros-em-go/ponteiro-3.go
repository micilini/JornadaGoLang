package main

import "fmt"

type Produto struct {
    Nome  string
    Preco float64
}

func aplicarDesconto(p *Produto) {
    p.Preco -= 10
}

func main() {
    p := Produto{"Notebook", 3500}
    aplicarDesconto(&p)
    fmt.Println("Preço com desconto:", p.Preco) // Saída: 3490
}