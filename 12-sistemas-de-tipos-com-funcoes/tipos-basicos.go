package main

import "fmt"

func soma(a int, b int) int {
    return a + b
}

func main() {
    resultado := soma(10, 5)
    fmt.Println("Soma:", resultado)
}