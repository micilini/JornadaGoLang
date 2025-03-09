package main

import "fmt"

func main() {
    var coisa interface{}

    // Atribuindo um inteiro
    coisa = 3
    fmt.Println(coisa) // Output: 3

    // Atribuindo uma string
    coisa = "olá, Go!"
    fmt.Println(coisa) // Output: olá, Go!

    // Atribuindo um float
    coisa = 3.14
    fmt.Println(coisa) // Output: 3.14

    // Atribuindo um booleano
    coisa = true
    fmt.Println(coisa) // Output: true
}
