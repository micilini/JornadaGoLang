package main

import "fmt"

func Soma[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Soma(10, 20))       // Inteiros
    fmt.Println(Soma(5.5, 2.3))     // Floats
}