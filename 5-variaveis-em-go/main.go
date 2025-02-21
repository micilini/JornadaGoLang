package main

var nomeGlobal string
var idadeGlobal int = 12
var tokenGlobal = "234123"

func main(){
	//Tipo Inferência:
	nome := "Lucas Galdino"
	idade := 25
	altura := 1.87

	//Declaração Explítica:
	var sobrenome string = "Roll"
	var salario int = 700
	var pi float64 = 3.1415

	//Declaração Tárdia:
	var laptop string
	var preco int

	laptop = "Dell 5440"
	preco = 6850

	//Declaração Simultânea:
	var nome2, sobrenome2 string = "Lucas", "Silva"
    var idade2, altura2 = 25, 1.75

	nome3, sobrenome3 := "Micilini", "Roll"
	idade3, altura3 := 22, 1.99

	//Agrupação por Blocos:
	var (
        nome4   = "Lucas"
        idade4  = 25
        altura4 = 1.75
        ativo4  = true
    )

	//Atualização de Variáveis
	var notebook = "Galaxy Book 2"

	notebook = "Galaxy Book 4 Pro"
	notebook = "Galaxy Book 5"
	notebook = 99
}