package main

import "fmt"

// Interface Forma define um comportamento comum.
type Forma interface {
    Area() float64
}

// Estrutura Quadrado
type Quadrado struct {
    Lado float64
}

// Método Area() para Quadrado (implementa Forma)
func (q Quadrado) Area() float64 {
    return q.Lado * q.Lado
}

// Estrutura Circulo
type Circulo struct {
    Raio float64
}

// Método Area() para Circulo (implementa Forma)
func (c Circulo) Area() float64 {
    return 3.14 * c.Raio * c.Raio
}

// NovaForma retorna uma interface Forma com base no tipo.
func NovaForma(tipo string, valor float64) Forma {
    if tipo == "quadrado" {
        return Quadrado{Lado: valor}
    }
    return Circulo{Raio: valor}
}

func main() {
    forma := NovaForma("quadrado", 5)
    fmt.Println("Área da forma:", forma.Area())

    outraForma := NovaForma("circulo", 3)
    fmt.Println("Área da outra forma:", outraForma.Area())
}