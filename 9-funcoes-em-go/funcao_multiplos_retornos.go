package main

import "fmt"

// Função que retorna múltiplos valores
func dividir(dividendo, divisor int) (int, int) {
    quociente := dividendo / divisor
    resto := dividendo % divisor
    return quociente, resto
}

func main() {
    q, r := dividir(10, 3)
    fmt.Printf("Quociente: %d, Resto: %d\n", q, r)
}
