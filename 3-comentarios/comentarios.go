// Este documento tem o objetivo de trazer todos os comentários suportados em GoLang

// Este é um comentário de exemplo!

// Este é um comentário de uma única linha

/*
   Este é um comentário de múltiplas linhas.
   Ele pode ser usado para explicar um bloco maior de código.
*/

// Este é um comentário do arquivo main.go
//  - vai mostrar uma mensagem de boas vindas
//   - Vai terminar a execução do programa

//go:build linux
package main

//Aqui estamos fazendo a importação dos nossos módulos...
import (
	"fmt"
)

func main() {
	/*
          Função principal da minha aplicação, aqui iremos mostrar a mensagem de boas vindas...
          A mensagem será "Olá Mundo!"
        */

        fmt.Println("Olá, Mundo!");
}

//Daqui pra baixo não tem mais nada...

// Pacote calculadora fornece funções matemáticas básicas.
package calculadora

// Soma retorna a soma de dois números inteiros.
func Soma(a, b int) int {
    return a + b
}

//go:noinline

func exemploFuncao() {
    // Essa função não será inline pelo compilador
}