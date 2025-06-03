package main

import (
    "fmt"
    "time"
)

func main() {
	locSP, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	tSP := time.Date(2025, 6, 2, 12, 0, 0, 0, locSP)
	fmt.Println("Horário SP:", tSP.Format(time.RFC3339))
}
