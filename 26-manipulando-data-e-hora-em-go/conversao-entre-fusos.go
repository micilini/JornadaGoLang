package main

import (
    "fmt"
    "time"
)

func main() {
	// t1 em UTC
	t1 := time.Now().UTC()

	// converter para horário de São Paulo
	locSP, _ := time.LoadLocation("America/Sao_Paulo")
	tSP := t1.In(locSP)
	fmt.Println("Em SP:", tSP)
}
