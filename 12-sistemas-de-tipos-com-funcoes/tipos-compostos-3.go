package main

import "fmt"

func contarPalavras(texto string) map[rune]int {
    contagem := make(map[rune]int)
    for _, letra := range texto {
        contagem[letra]++
    }
    return contagem
}

func main() {
    resultado := contarPalavras("banana")
    fmt.Println("Frequência das letras:", resultado)
}
