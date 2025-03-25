package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Estrutura genérica para armazenar qualquer elemento XML
type Elemento struct {
	XMLName xml.Name            `xml:""`
	Atributos map[string]string `xml:",any,attr"`
	Conteudo  string            `xml:",chardata"`
}

func main() {
	// Abrir o arquivo XML desconhecido
	arquivo := "produto.xml"
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um decoder XML
	decoder := xml.NewDecoder(file)

	// Ler o XML de forma genérica
	var elemento Elemento
	if err := decoder.Decode(&elemento); err != nil {
		fmt.Println("Erro ao decodificar XML:", err)
		return
	}

	// Exibir os dados lidos
	fmt.Printf("Elemento: %s\n", elemento.XMLName.Local)
	fmt.Println("Atributos:")
	for k, v := range elemento.Atributos {
		fmt.Printf("  %s = %s\n", k, v)
	}
	fmt.Println("Conteúdo:", elemento.Conteudo)
}