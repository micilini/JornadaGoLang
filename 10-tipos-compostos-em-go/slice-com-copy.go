package main

import "fmt"

func main() {
    origem := []int{1, 2, 3, 4, 5}
    destino := make([]int, len(origem))

    n := copy(destino, origem)

    fmt.Println("Origem:", origem)     // [1 2 3 4 5]
    fmt.Println("Destino:", destino)   // [1 2 3 4 5]
    fmt.Println("Elementos copiados:", n) // 5
}