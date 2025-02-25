package main

import "fmt"

func main() {
    contador := 0

    for contador < 5 {
        fmt.Println("Contador:", contador)
        contador++
    }

	for {
        fmt.Println("Executando para sempre...")
    }

	frutas := []string{"maçã", "banana", "uva"}

    for indice, fruta := range frutas {
        fmt.Printf("Índice: %d, Fruta: %s\n", indice, fruta)
    }

	for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("Parando no 5")
            break
        }
        fmt.Println(i)
    }

	for i := 0; i < 5; i++ {
        if i == 2 {
            continue // Pula o número 2
        }
        fmt.Println(i)
    }
}
