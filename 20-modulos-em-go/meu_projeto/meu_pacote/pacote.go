package meu_pacote

import "fmt"

// Exportado (público)
func MinhaFuncao() {
    fmt.Println("Função Exportada!")
}

// Privado ao pacote
func minhaFuncaoPrivada() {
    fmt.Println("Apenas para este pacote!")
}