package main

import "fmt"

func main() {
    dia := "segunda"

    switch dia {
    case "segunda":
        fmt.Println("Início da semana!")
    case "sexta":
        fmt.Println("Quase final de semana!")
    case "domingo":
        fmt.Println("Dia de descanso!")
    default:
        fmt.Println("Dia comum da semana.")
    }

    // Switch com múltiplos valores
    switch dia {
    case "sábado", "domingo":
        fmt.Println("É fim de semana!")
    default:
        fmt.Println("Dia útil.")
    }

    // Switch sem expressão (Switch True)
    idade := 25
    switch {
    case idade < 18:
        fmt.Println("Menor de idade.")
    case idade >= 18 && idade < 60:
        fmt.Println("Adulto.")
    default:
        fmt.Println("Idoso.")
    }

    // Usando fallthrough
    nota := 8
    switch {
    case nota >= 9:
        fmt.Println("Excelente!")
        fallthrough
    case nota >= 7:
        fmt.Println("Aprovado!")
    default:
        fmt.Println("Reprovado.")
    }

    // Type Switch (Switch com tipos)
    identificar(42)
    identificar("GoLang")
    identificar(true)
}

func identificar(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("É um inteiro:", v)
    case string:
        fmt.Println("É uma string:", v)
    case bool:
        fmt.Println("É um booleano:", v)
    default:
        fmt.Println("Tipo desconhecido")
    }
}