package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	arquivo := "produto-2.xml"
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	// Percorrer os tokens do XML
	for {
		token, err := decoder.Token()
		if err != nil {
			break // Fim do arquivo
		}

		// Identificar o tipo do token (elemento de abertura, fechamento, conteúdo etc.)
		switch elemento := token.(type) {
		case xml.StartElement:
			fmt.Printf("Início do elemento: %s\n", elemento.Name.Local)
			for _, attr := range elemento.Attr {
				fmt.Printf("  Atributo: %s = %s\n", attr.Name.Local, attr.Value)
			}
		case xml.CharData:
			conteudo := string(elemento)
			if len(conteudo) > 0 {
				fmt.Println("  Conteúdo:", conteudo)
			}
		case xml.EndElement:
			fmt.Printf("Fim do elemento: %s\n", elemento.Name.Local)
		}
	}
}