package main

import "fmt"

func main() {
    var meuArray [10]string

    for i := 0; i < len(meuArray); i++ {
        meuArray[i] = "Fruta"
    }

    fmt.Println(meuArray) // Saída: [Fruta Fruta Fruta ...]
}