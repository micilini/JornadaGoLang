package main

import (
    "fmt"
    "time"
)

func main() {
	// 2 horas + 45 minutos + 30 segundos
	t := time.Now()
	d := 2*time.Hour + 45*time.Minute + 30*time.Second
	tNovo := t.Add(d)
	fmt.Println("Depois de 2h45m30s:", tNovo)
}
