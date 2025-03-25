package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Estrutura conhecida
type Loja struct {
	XMLName  xml.Name  `xml:"loja"`
	Produtos []Produto `xml:"produto"`
}

type Produto struct {
	XMLName xml.Name `xml:"produto"`
	ID      string   `xml:"id,attr"`
	Nome    string   `xml:"nome"`
	Preco   Preco    `xml:"preco"`
}

type Preco struct {
	XMLName xml.Name `xml:"preco"`
	Moeda   string   `xml:"moeda,attr"`
	Valor   string   `xml:",chardata"`
}

func main() {
	arquivo := "produto-2.xml"
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var loja Loja
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&loja); err != nil {
		fmt.Println("Erro ao decodificar XML:", err)
		return
	}

	// Exibir produtos extraídos
	fmt.Println("Produtos extraídos:")
	for _, p := range loja.Produtos {
		fmt.Printf("ID: %s, Nome: %s, Preço: %s %s\n", p.ID, p.Nome, p.Preco.Moeda, p.Preco.Valor)
	}
}