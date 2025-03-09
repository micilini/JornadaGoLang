package main

import "fmt"

func main() {
    var meuArray [10]string

    // Atribuindo valores por índice
    meuArray[0] = "Maçã"
    meuArray[1] = "Banana"
    meuArray[9] = "Uva"

    fmt.Println(meuArray) // Saída: [Maçã Banana       Uva]
}