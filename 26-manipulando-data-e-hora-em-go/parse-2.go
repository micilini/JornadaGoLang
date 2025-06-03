package main

import (
    "fmt"
    "time"
)

func main() {
	// Exemplo 1: converter "2025-06-15" para time.Time
	layout := "2006-01-02"
	strData := "2025-06-15"
	t, err := time.Parse(layout, strData)
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return
	}
	fmt.Println("Data convertida:", t) // Resultado em UTC: 2025-06-15 00:00:00 +0000 UTC
}