package main

import (
    "fmt"
    "time"
)

func main() {
	// Exemplo com fuso horário de São Paulo
	locSP, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	tSP := time.Date(2025, time.June, 15, 14, 30, 0, 0, locSP)
	fmt.Println("Data em horário de SP:", tSP)
}
