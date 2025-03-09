package main

import "fmt"

func main() {
	numeros := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

    // Percorrendo o array com for range
    for indice, valor := range numeros {
        fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
    }
}