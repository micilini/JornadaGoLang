package main

import "fmt"

func main() {
    a := 10
    b := 20

    fmt.Println("a == b:", a == b)  // false (10 não é igual a 20)
    fmt.Println("a != b:", a != b)  // true (10 é diferente de 20)
    fmt.Println("a < b:", a < b)    // true (10 é menor que 20)
    fmt.Println("a <= b:", a <= b)  // true (10 é menor ou igual a 20)
    fmt.Println("a > b:", a > b)    // false (10 não é maior que 20)
    fmt.Println("a >= b:", a >= b)  // false (10 não é maior ou igual a 20)

    // Comparando strings
    nome1 := "Go"
    nome2 := "Golang"

    fmt.Println("nome1 == nome2:", nome1 == nome2) // false
    fmt.Println("nome1 != nome2:", nome1 != nome2) // true
}