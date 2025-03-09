package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p1 := Pessoa{"João", 25}
	p2 := Pessoa{"João", 25}
        p3 := Pessoa{"Micilini", 250}

	fmt.Println(p1 == p2) // Saída: true
        fmt.Println(p1 == p3) // Saída: false
}
