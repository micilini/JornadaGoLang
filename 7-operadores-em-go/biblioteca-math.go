package main

import (
    "fmt"
    "math"
)

func main() {
    // Operações básicas
    fmt.Println("Abs(-7.5):", math.Abs(-7.5))
    fmt.Println("Max(3, 7):", math.Max(3, 7))
    fmt.Println("Min(3, 7):", math.Min(3, 7))
    fmt.Println("Mod(10, 3):", math.Mod(10, 3))

    // Potências e raízes
    fmt.Println("Pow(2, 3):", math.Pow(2, 3))
    fmt.Println("Sqrt(25):", math.Sqrt(25))
    fmt.Println("Cbrt(27):", math.Cbrt(27))

    // Funções trigonométricas
    fmt.Println("Sin(Pi/2):", math.Sin(math.Pi/2))
    fmt.Println("Cos(0):", math.Cos(0))
    fmt.Println("Tan(Pi/4):", math.Tan(math.Pi/4))

    // Arredondamento
    fmt.Println("Floor(3.7):", math.Floor(3.7))
    fmt.Println("Ceil(3.1):", math.Ceil(3.1))
    fmt.Println("Round(3.5):", math.Round(3.5))
    fmt.Println("Trunc(3.9):", math.Trunc(3.9))

    // Logaritmos
    fmt.Println("Log(10):", math.Log(10))
    fmt.Println("Log2(8):", math.Log2(8))
    fmt.Println("Log10(100):", math.Log10(100))
}