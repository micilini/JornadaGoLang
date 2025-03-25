package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Definição das estruturas XML
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

	// Exibir os dados antes da adição
	fmt.Println("Antes da adição:")
	for _, pessoa := range pessoas.Lista {
		fmt.Printf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	}

	// Criar um novo registro
	novoRegistro := Pessoa{
		Nome:      "Carlos",
		Idade:     40,
		Profissao: "Analista",
	}

	// Adicionar o novo registro à lista existente
	pessoas.Lista = append(pessoas.Lista, novoRegistro)

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

	// Exibir os dados após a adição
	fmt.Println("\nDepois da adição:")
	for _, pessoa := range pessoas.Lista {
		fmt.Printf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	}

	fmt.Println("\nNovo registro adicionado e arquivo atualizado com sucesso!")
}