package main

import "fmt"

// Função genérica que aceita um parâmetro de tipo T
func imprimir[T any](valor T) {
    fmt.Println(valor)
}

func main() {
    imprimir(42)            // Tipo int
    imprimir("Olá, Go!")    // Tipo string
    imprimir(3.14)          // Tipo float
    imprimir(true)          // Tipo bool
}