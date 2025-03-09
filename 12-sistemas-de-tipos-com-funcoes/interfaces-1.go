package main

import "fmt"

type Forma interface {
    Area() float64
}

type Quadrado struct {
    Lado float64
}

func (q Quadrado) Area() float64 {
    return q.Lado * q.Lado
}

type Circulo struct {
    Raio float64
}

func (c Circulo) Area() float64 {
    return 3.14 * c.Raio * c.Raio
}

func imprimirArea(f Forma) {
    fmt.Println("Área da forma:", f.Area())
}

func main() {
    q := Quadrado{Lado: 4}
    c := Circulo{Raio: 3}

    imprimirArea(q)
    imprimirArea(c)
}