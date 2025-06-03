package main

import (
    "fmt"
    "time"
)

func main() {
	t:= time.Now()

	// Adiciona 1 ano, 2 meses e 10 dias
	tFuturo := t.AddDate(1, 2, 10)

	// Subtrai 0 anos, 3 meses e 5 dias
	tPassado := t.AddDate(0, -3, -5)

	fmt.Println("Agora:", t)
	fmt.Println("Futuro (+1a+2m+10d):", tFuturo)
	fmt.Println("Passado (-3m-5d):", tPassado)
}
