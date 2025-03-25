package main

import (
    "fmt"
    "meu_projeto/meu_pacote" // Importando o pacote
)

func main() {
    meu_pacote.MinhaFuncao() // OK: Função pública
    // meu_pacote.minhaFuncaoPrivada() // Erro: Não exportada

    fmt.Println("Executando main.go!")
}