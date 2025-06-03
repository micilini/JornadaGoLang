package main

import (
    "fmt"
    "time"
)

func main() {
	inicio := time.Now()
	// ...faz algo demorado...
	duracao := time.Since(inicio)
	fmt.Println("Levantamento levou:", duracao) // ex: "2.3456789s"
}