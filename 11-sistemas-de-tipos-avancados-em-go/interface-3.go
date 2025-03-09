package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var s imprimivel = pessoa{"Micilini", "Roll"}
	fmt.Println(s.toString())
	imprimir(s)

	p = produto{"Oculus Rift 3", 5800.89}
	fmt.Println(p.toString())
	imprimir(p)

	p2 := produto{"RTX 59090 TI", 30000.99}
	imprimir(p2)
}