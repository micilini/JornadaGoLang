package main

import (
    "fmt"
    "time"
)

func main() {
	t1 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)

	d := t2.Sub(t1) // time.Duration representando 14 dias e 12 horas
	fmt.Println("Diferença em horas:", d.Hours())
	fmt.Println("Diferença em dias (aprox):", d.Hours()/24)
}