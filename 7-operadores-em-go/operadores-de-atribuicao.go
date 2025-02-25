package main

import "fmt"

func main() {
    var x int = 10

    // Atribuição simples
    fmt.Println("Valor inicial de x:", x)

    // Adição e atribuição
    x += 5
    fmt.Println("x += 5 →", x) // 15

    // Subtração e atribuição
    x -= 3
    fmt.Println("x -= 3 →", x) // 12

    // Multiplicação e atribuição
    x *= 2
    fmt.Println("x *= 2 →", x) // 24

    // Divisão e atribuição
    x /= 4
    fmt.Println("x /= 4 →", x) // 6

    // Módulo e atribuição
    x %= 5
    fmt.Println("x %= 5 →", x) // 1

    // Operações bit a bit
    x = 10
    x &= 6
    fmt.Println("x &= 6 →", x) // 2

    x |= 3
    fmt.Println("x |= 3 →", x) // 3

    x ^= 1
    fmt.Println("x ^= 1 →", x) // 2

    // Deslocamento de bits
    x <<= 2
    fmt.Println("x <<= 2 →", x) // 8

    x >>= 1
    fmt.Println("x >>= 1 →", x) // 4
}