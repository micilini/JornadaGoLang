package main

import "fmt"

func main() {
    a := true
    b := false

    // Operador AND (&&) - Ambas as condições devem ser verdadeiras
    fmt.Println("a && b:", a && b) // false (true E false → false)
    fmt.Println("true && true:", true && true) // true

    // Operador OR (||) - Pelo menos uma condição deve ser verdadeira
    fmt.Println("a || b:", a || b) // true (true OU false → true)
    fmt.Println("false || false:", false || false) // false

    // Operador NOT (!) - Inverte o valor booleano
    fmt.Println("!a:", !a) // false (inverte true → false)
    fmt.Println("!b:", !b) // true (inverte false → true)
}