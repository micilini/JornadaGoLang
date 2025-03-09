package main

import "fmt"

type Contador struct {
    Valor int
}

func (c *Contador) Incrementar() {
    c.Valor++
}

func main() {
    c := Contador{Valor: 0}
    c.Incrementar()
    fmt.Println("Valor do contador:", c.Valor)
}