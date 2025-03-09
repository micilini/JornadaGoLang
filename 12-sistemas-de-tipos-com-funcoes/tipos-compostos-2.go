package main

import "fmt"

// RetornaSlice recebe dois números e os retorna em um slice.
func RetornaSlice(a, b int) []int {
    return []int{a, b}
}

func main() {
    numeros := RetornaSlice(10, 20)
    fmt.Println("Slice retornado:", numeros)
}