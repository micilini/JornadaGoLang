package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

	// 2️⃣ Modificar a linha desejada
	for i, linha := range dados {
		if linha[0] == "2" { // Procuramos a linha onde o ID é "2" (Bob)
			dados[i][2] = strconv.Itoa(26) // Atualizando a idade para 26
			break
		}
	}

	// 3️⃣ Escrever os dados de volta para o arquivo (sobrescrevendo)
	file, err = os.Create(arquivo) // Criamos um novo arquivo (sobrescrevendo)
	if err != nil {
		fmt.Println("Erro ao sobrescrever o arquivo:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(dados) // Escreve todas as linhas novamente
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	writer.Flush() // Garante que os dados sejam gravados

	fmt.Println("Linha atualizada com sucesso!")
}