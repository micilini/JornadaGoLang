package main

import "fmt"

func main() {
    x := 10
    p := &x // 'p' é um ponteiro para 'x'

    fmt.Println("Valor de x:", x)   // Saída: 10
    fmt.Println("Endereço de x:", p) // Saída: endereço de memória
    fmt.Println("Valor via ponteiro:", *p) // Saída: 10

    *p = 20 // Modifica o valor de 'x' através do ponteiro
    fmt.Println("Novo valor de x:", x) // Saída: 20
}
