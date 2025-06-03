package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Printf("Ano: %d\n", t.Year())
    fmt.Printf("Mês: %s (%d)\n", t.Month(), t.Month())
    fmt.Printf("Dia: %d\n", t.Day())
    fmt.Printf("Hora: %d\n", t.Hour())
    fmt.Printf("Minuto: %d\n", t.Minute())
    fmt.Printf("Segundo: %d\n", t.Second())
    fmt.Printf("Dia da semana: %s\n", t.Weekday())
}
