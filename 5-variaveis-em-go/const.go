package main

import "fmt"

// Constantes globais (declaradas fora de qualquer função)
const (
    Pi       = 3.14159
    Language = "Go"
)

// Constantes dentro da função main (constantes locais)
func main() {
    // Constantes com valor numérico
    const number = 10

    // Constantes de tipo específico
    const (
        Year int = 2025
        Month     = 2
    )

    // Constantes de inferência de tipo
    const (
        Name     = "John Doe"
        IsActive = true
    )

    // Exemplo de bloco de constantes dentro do main
    const (
        ServerName   = "localhost"
        ServerPort   = 8080
        MaxClients   = 100
    )

}