package main

import (
    "fmt"
    "time"
)

func checkVencimento(dataVencimento time.Time) {
    agora := time.Now()
    if agora.After(dataVencimento) {
        fmt.Println("A data de vencimento já passou!")
    } else if agora.Equal(dataVencimento) {
        fmt.Println("Hoje é o dia do vencimento!")
    } else {
        fmt.Println("Ainda falta tempo para vencimento.")
    }
}

func main() {
    venc := time.Date(2025, 6, 10, 0, 0, 0, 0, time.Local)
    checkVencimento(venc)
}
