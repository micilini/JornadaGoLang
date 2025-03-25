package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Definição da estrutura
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
	pessoas := Pessoas{
		Lista: []Pessoa{
			{Nome: "Alice", Idade: 30, Profissao: "Engenheira"},
			{Nome: "Bob", Idade: 25, Profissao: "Desenvolvedor"},
		},
	}

	file, err := os.Create("dados.xml")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(pessoas); err != nil {
		fmt.Println("Erro ao escrever XML:", err)
	}
	fmt.Println("Arquivo XML criado com sucesso!")
}