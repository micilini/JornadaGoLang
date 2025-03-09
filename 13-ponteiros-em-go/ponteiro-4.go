package main

import "fmt"

func getNome() *string {
    return nil // Não há nome disponível
}

func main() {
    nome := getNome()

    if nome == nil {
        fmt.Println("Nome não encontrado!")
    }
}