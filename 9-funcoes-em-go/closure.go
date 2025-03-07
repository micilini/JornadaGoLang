package main

import "fmt"

func contador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}

func main() {
    incrementar := contador()

    fmt.Println(incrementar()) // Saída: 1
    fmt.Println(incrementar()) // Saída: 2
    fmt.Println(incrementar()) // Saída: 3
}
