package main

import "fmt"

func main() {
    idade := 10

    if idade == 10 {
        fmt.Println("O valor é igual a 10")
    } else {
        fmt.Println("O valor não é igual a 10")
    }

	temperatura := 25

    if temperatura > 30 {
        fmt.Println("Está muito quente!")
    } else if temperatura > 20 {
        fmt.Println("Clima agradável.")
    } else {
        fmt.Println("Está frio.")
    }
}
