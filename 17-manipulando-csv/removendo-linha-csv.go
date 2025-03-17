package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo := "dados.csv"

	// 1️⃣ Abrir o arquivo para leitura
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	dados, err := reader.ReadAll() // Lê todo o conteúdo do CSV
	if err != nil {
		fmt.Println("Erro ao ler CSV:", err)
		return
	}

	// 2️⃣ Remover a linha desejada (ID "2" -> Bob)
	var novosDados [][]string
	for _, linha := range dados {
		if linha[0] != "2" { // Se o ID não for "2", mantemos a linha
			novosDados = append(novosDados, linha)
		}
	}

	// 3️⃣ Escrever os dados atualizados de volta para o arquivo
	file, err = os.Create(arquivo) // Criamos um novo arquivo (sobrescrevendo)
	if err != nil {
		fmt.Println("Erro ao sobrescrever o arquivo:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(novosDados) // Escreve os dados restantes (sem a linha removida)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	writer.Flush() // Garante que os dados sejam gravados

	fmt.Println("Linha removida com sucesso!")
}