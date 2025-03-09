package main

import "fmt"

// Função que aceita um valor de qualquer tipo
func imprimirValor(v any) {
    fmt.Println(v)
}

func main() {
    // Chamando a função com diferentes tipos
    imprimirValor(42)          // Um inteiro
    imprimirValor("Olá, Go!")  // Uma string
    imprimirValor(3.14)        // Um float
    imprimirValor(true)        // Um booleano
}
