package main

import "fmt"

func main() {
    // Criando um map com chave string e valor int
    idades := map[string]int{
        "Alice": 30,
        "Bob":   25,
        "Eve":   22,
    }

    fmt.Println(idades) // Saída: map[Alice:30 Bob:25 Eve:22]

	contagem := make(map[string]int) // Cria um map vazio
	contagem["Maçã"] = 5
	contagem["Banana"] = 3

	fmt.Println(contagem) // Saída: map[Maçã:5 Banana:3]

	pontos := make(map[string]int)

	pontos["João"] = 50 // Adiciona um novo par
	pontos["Ana"] = 75  // Adiciona outro par
	pontos["João"] = 100 // Atualiza um valor existente

	fmt.Println(pontos) // Saída: map[Ana:75 João:100]

	notas := map[string]float64{
		"Matemática": 9.5,
		"História":   8.0,
	}
	
	fmt.Println(notas["Matemática"]) // Saída: 9.5

	nomes := map[string]string{
		"123": "Carlos",
		"456": "Ana",
	}
	
	delete(nomes, "123") // Remove o item com a chave "123"
	
	fmt.Println(nomes) // Saída: map[456:Ana]

	idades := map[string]int{"João": 30}

	idade, existe := idades["Ana"]

	if existe {
		fmt.Println("Idade encontrada:", idade)
	} else {
		fmt.Println("Chave não encontrada")
	}

	frutas := map[string]int{
		"Maçã":    5,
		"Banana":  3,
		"Laranja": 2,
	}
	
	for chave, valor := range frutas {
		fmt.Printf("%s: %d\n", chave, valor)
	}

	meuMapa = make(map[string]int) // Novo map vazio

	original := map[string]int{"A": 1}
	copia := original

	copia["A"] = 42
	fmt.Println(original) // Saída: map[A:42]

	novo := make(map[string]int)
	for k, v := range original {
		novo[k] = v
	}
}