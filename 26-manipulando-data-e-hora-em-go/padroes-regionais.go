package main

import (
    "fmt"
    "time"
)

func main() {
	meses := []string{
		"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	}

	t := time.Now()
	dia := t.Day()
	mesExtenso := meses[int(t.Month())-1]
	ano := t.Year()

	fmt.Printf("%02d de %s de %d\n", dia, mesExtenso, ano)
	// Ex.: "02 de junho de 2025"
}