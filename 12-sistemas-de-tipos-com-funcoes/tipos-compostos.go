package main

import "fmt"

func somaElementos(numeros []int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    fmt.Println("Soma dos elementos:", somaElementos(numeros))
}
