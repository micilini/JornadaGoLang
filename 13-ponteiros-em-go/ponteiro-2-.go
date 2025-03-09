package main

import "fmt"

func dobrarValor(n *int) {
    *n = *n * 2
}

func main() {
    num := 10
    dobrarValor(&num)
    fmt.Println("Com ponteiro:", num) // Saída: 20 (mudou!)
}