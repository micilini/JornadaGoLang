package main

import "fmt"

// Função com parâmetro
func saudacao(nome string) {
    fmt.Println("Olá, " + nome)
}

func main() {
    saudacao("Micilini")
    saudacao("Roll")
}
