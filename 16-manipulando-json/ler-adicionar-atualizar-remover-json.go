package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func lerArquivoJSON(caminho string) ([]map[string]interface{}, error) {
	file, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pessoas []map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pessoas)
	if err != nil {
		return nil, err
	}
	return pessoas, nil
}

func adicionarPessoa(pessoas *[]map[string]interface{}, nome string, idade int, profissao string) {
	novaPessoa := map[string]interface{}{
		"nome":  nome,
		"idade": idade,
	}
	if profissao != "" {
		novaPessoa["profissao"] = profissao
	}
	*pessoas = append(*pessoas, novaPessoa)
}

func atualizarPessoa(pessoas *[]map[string]interface{}, nome string, novaIdade int, novaProfissao string) bool {
	for i, pessoa := range *pessoas {
		if pessoa["nome"] == nome {
			(*pessoas)[i]["idade"] = novaIdade
			if novaProfissao != "" {
				(*pessoas)[i]["profissao"] = novaProfissao
			}
			return true
		}
	}
	return false
}

func removerPessoa(pessoas *[]map[string]interface{}, nome string) bool {
	for i, pessoa := range *pessoas {
		if pessoa["nome"] == nome {
			*pessoas = append((*pessoas)[:i], (*pessoas)[i+1:]...)
			return true
		}
	}
	return false
}

func salvarArquivoJSON(caminho string, pessoas []map[string]interface{}) error {
	file, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Indentação para melhor visualização
	return encoder.Encode(pessoas)
}

func main() {
	// Caminho do arquivo JSON
	caminho := "pessoas.json"

	// Ler o arquivo JSON
	pessoas, err := lerArquivoJSON(caminho)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}

	// Adicionar uma nova pessoa
	adicionarPessoa(&pessoas, "Lucas", 27, "Designer")

	// Atualizar uma pessoa
	atualizarPessoa(&pessoas, "Carlos", 41, "Gerente de Engenharia")

	// Remover uma pessoa
	removerPessoa(&pessoas, "Maria")

	// Salvar alterações no arquivo
	err = salvarArquivoJSON(caminho, pessoas)
	if err != nil {
		fmt.Println("Erro ao salvar JSON:", err)
		return
	}

	fmt.Println("Alterações salvas com sucesso!")
}