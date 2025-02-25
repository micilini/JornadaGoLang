package main

import "fmt"

func main() {
    fmt.Println("Início")

    // Salta para o rótulo "Fim"
    goto Fim

    fmt.Println("Isso não será impresso.")

Fim:
    fmt.Println("Fim do programa.")

    // Exemplo com erro detectado
    for i := 1; i <= 5; i++ {
        if i == 3 {
            goto Erro
        }
        fmt.Println("Processando:", i)
    }

Erro:
    fmt.Println("Erro detectado! Saindo do loop.")

    // Usando goto para sair de loops aninhados
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i=%d, j=%d\n", i, j)
            if i == 2 && j == 2 {
                goto Saida
            }
        }
    }

Saida:
    fmt.Println("Loop encerrado.")
}
