package main

import "fmt"

// Função que retorna um valor inteiro
func soma(a int, b int) int {
    return a + b
}

func main() {
    resultado := soma(5, 7)
    fmt.Println("Resultado da soma:", resultado)
}
