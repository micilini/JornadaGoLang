package main

import "fmt"

func main() {
    // 1. Operador de Sinal (+ e -)
    a := 10
    b := -a
    fmt.Println("+a:", +a) // Saída: 10
    fmt.Println("-a:", b)  // Saída: -10

    // 2. Operador de Negação Lógica (!)
    ativo := true
    fmt.Println("!ativo:", !ativo) // Saída: false

    // 3. Operador de Endereço (&) e Desreferência (*)
    x := 42
    ptr := &x       // & obtém o endereço de memória de x
    fmt.Println("Endereço de x:", ptr)

    fmt.Println("Valor apontado por ptr:", *ptr) // * acessa o valor no endereço

    // 4. Operador de Complemento Bit a Bit (^)
    num := 5        // Em binário: 101
    fmt.Println("^num:", ^num) // Inverte os bits: -6 (em complemento de dois)
}