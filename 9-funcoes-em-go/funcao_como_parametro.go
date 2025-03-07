package main

import "fmt"

// Função que aceita outra função como parâmetro
func executar(fn func(int, int) int, a int, b int) {
    resultado := fn(a, b)
    fmt.Println("Resultado:", resultado)
}

func main() {
    soma := func(x, y int) int { return x + y }
    executar(soma, 3, 4)
}
