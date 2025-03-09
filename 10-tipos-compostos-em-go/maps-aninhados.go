package main

import "fmt"

func main() {
    estrutura := map[string]map[string]map[string]int{
        "primeiro": {
            "segundo": {
                "terceiro": 42,
            },
        },
    }

    fmt.Println(estrutura["primeiro"]["segundo"]["terceiro"]) // Saída: 42
}