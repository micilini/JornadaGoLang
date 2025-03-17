package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("dados.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Eve", "35", "Arquiteta"})

	fmt.Println("Nova linha adicionada ao CSV!")
}