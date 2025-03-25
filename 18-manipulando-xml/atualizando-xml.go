package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Definição das estruturas para o XML
type Pessoa struct {
	XMLName   xml.Name `xml:"pessoa"`
	Nome      string   `xml:"nome"`
	Idade     int      `xml:"idade"`
	Profissao string   `xml:"profissao"`
}

type Pessoas struct {
	XMLName xml.Name `xml:"pessoas"`
	Lista   []Pessoa `xml:"pessoa"`
}

func main() {
	// Nome do arquivo XML
	arquivo := "dados.xml"

	// Abrir o arquivo XML para leitura
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um decoder XML
	decoder := xml.NewDecoder(file)

	// Criar a variável para armazenar os dados
	var pessoas Pessoas

	// Decodificar o XML para a struct
	if err := decoder.Decode(&pessoas); err != nil {
		fmt.Println("Erro ao decodificar XML:", err)
		return
	}

	// Exibir os dados antes da alteração
	fmt.Println("Antes da alteração:")
	for _, pessoa := range pessoas.Lista {
		fmt.Printf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	}

	// Alterar a idade da pessoa chamada "Alice"
	for i := range pessoas.Lista {
		if pessoas.Lista[i].Nome == "Alice" {
			pessoas.Lista[i].Idade = 35 // Nova idade
		}
	}

	// Criar um novo arquivo para salvar as alterações
	file, err = os.Create(arquivo)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um encoder XML para salvar os dados
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ") // Formatação bonita

	// Escrever os dados atualizados no XML
	if err := encoder.Encode(pessoas); err != nil {
		fmt.Println("Erro ao escrever XML:", err)
		return
	}

	// Exibir os dados após a alteração
	fmt.Println("\nDepois da alteração:")
	for _, pessoa := range pessoas.Lista {
		fmt.Printf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	}

	fmt.Println("\nArquivo atualizado com sucesso!")
}