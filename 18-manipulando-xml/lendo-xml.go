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
	// Abrir o arquivo XML
	file, err := os.Open("dados.xml")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um decoder XML
	decoder := xml.NewDecoder(file)

	// Criar uma variável para armazenar os dados
	var pessoas Pessoas

	// Decodificar o XML para a struct
	if err := decoder.Decode(&pessoas); err != nil {
		fmt.Println("Erro ao decodificar XML:", err)
		return
	}

	// Exibir os dados lidos
	for _, pessoa := range pessoas.Lista {
		fmt.Printf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	}
}