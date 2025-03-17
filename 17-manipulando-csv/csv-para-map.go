package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Abrir o arquivo CSV
	file, err := os.Open("dados.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um leitor CSV
	reader := csv.NewReader(file)

	// Ler todo o conteúdo do arquivo CSV para um slice de slices
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler arquivo CSV:", err)
		return
	}

	// Criar um mapa para armazenar os dados por coluna
	dados := make(map[string][]string)

	// Ler as colunas (cabeçalho)
	cabecalho := records[0]

	// Inicializar o mapa com as chaves vazias
	for _, coluna := range cabecalho {
		dados[coluna] = []string{}
	}

	// Preencher os dados a partir das linhas
	for _, record := range records[1:] {
		for i, valor := range record {
			dados[cabecalho[i]] = append(dados[cabecalho[i]], valor)
		}
	}

	// Exibir o mapa com os dados por coluna
	for chave, valores := range dados {
		fmt.Println(chave, ":", valores)
	}
}