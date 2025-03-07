package main

import "fmt"

// Função variádica
func somar(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

func main() {
    fmt.Println(somar(1, 2, 3, 4, 5)) // Saída: 15
}
