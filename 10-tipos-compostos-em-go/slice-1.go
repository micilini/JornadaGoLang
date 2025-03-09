package main

import "fmt"

func main() {
	frutasUm := []string{"maçã", "banana", "uva"}

	var frutasDois []string

	frutas := make([]string, 0, 10)

	fmt.Println(frutasUm[0]) // Imprime o primeiro elemento, "maçã"
	frutas = append(frutas, "laranja") // Adiciona "laranja" à slice
	frutasSubset := frutas[1:3] // Cria uma slice com os elementos de índice 1 e 2

	for _, fruta := range frutas {
        fmt.Printf("Fruta: %s\n", i, fruta)
	}

	nums := []int{1, 2, 3}
	nums = append(nums, 4)  // nums agora é [1, 2, 3, 4]

	s := []int{1, 2, 3}
    fmt.Println(len(s), cap(s)) // Inicialmente len = 3, cap = 3

    s = append(s, 4) // Adiciona um elemento, a capacidade pode ser duplicada
    fmt.Println(len(s), cap(s)) // Agora len = 4, cap = 6 (ou algo maior, dependendo da implementação)

	// Criando uma slice de inteiros com comprimento 3 e capacidade 5
    s2 := make([]int, 3, 5)

    fmt.Println(s2)         // Imprime: [0 0 0]
    fmt.Println(len(s2))    // Imprime: 3 (comprimento)
    fmt.Println(cap(s2))    // Imprime: 5 (capacidade)
}
