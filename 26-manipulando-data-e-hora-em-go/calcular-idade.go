package main

import (
    "fmt"
    "time"
)

func calcularIdade(nascimento time.Time) int {
    hoje := time.Now()
    idade := hoje.Year() - nascimento.Year()

    // Se ainda não chegou ao mês/dia de aniversário neste ano, subtrai 1
    if hoje.Month() < nascimento.Month() ||
        (hoje.Month() == nascimento.Month() && hoje.Day() < nascimento.Day()) {
        idade--
    }
    return idade
}

func main() {
    // Supondo data de nascimento: 15 de agosto de 1990
    nascimento := time.Date(1990, time.August, 15, 0, 0, 0, 0, time.UTC)
    fmt.Println("Idade:", calcularIdade(nascimento)) // Ex.: 34 (em jun/2025)
}
