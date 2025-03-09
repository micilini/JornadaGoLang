package main

import "fmt"

func main() {
    frutas := []string{"maçã", "banana", "uva", "laranja", "morango"}

    // Fatiamento até o índice 3 (não inclui o elemento no índice 3)
    parte := frutas[:3]

    fmt.Println(parte) // Imprime: [maçã banana uva]
}
