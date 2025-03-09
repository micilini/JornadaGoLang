package main

import "fmt"

func main() {
    // Slice que armazena diferentes tipos de dados
    mixedSlice := []interface{}{42, "texto", 3.14, true}

    for _, v := range mixedSlice {
        fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
    }
}
