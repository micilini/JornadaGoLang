package main

import "fmt"

// Função genérica com múltiplos parâmetros de tipo
func somar[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(somar(10, 20))       // Soma com int
    fmt.Println(somar(10.5, 20.5))   // Soma com float64
}