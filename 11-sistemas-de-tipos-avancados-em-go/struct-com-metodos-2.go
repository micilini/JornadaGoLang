package main

import "fmt"

// Definindo a struct Endereco
type Endereco struct {
    Rua    string
    Numero int
    Cidade string
}

// Função que recebe um Endereco e exibe as informações no console
func mostrarEndereco(endereco Endereco) {
    fmt.Println("Endereço completo:")
    fmt.Println("Rua:", endereco.Rua)
    fmt.Println("Número:", endereco.Numero)
    fmt.Println("Cidade:", endereco.Cidade)
}

func main() {
    // Criando uma instância de Endereco
    meuEndereco := Endereco{
        Rua:    "Rua das Palmeiras",
        Numero: 123,
        Cidade: "São Paulo",
    }

    // Chamando a função passando o Endereco
    mostrarEndereco(meuEndereco)
}