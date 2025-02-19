package main

import "fmt"

func main(){
	x := 12
	fmt.Println("Olá " + "Mundo! "  + x)

	//Print():
	fmt.Print("Olá, ")
	fmt.Print("Micilini!")

	//Println():
	fmt.Println("Vamos aprender: ")
	fmt.Println("GoLANG!")

	//Printf():
	nome := "João Vitor"
    idade := 25
    fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)

	//Scan():
	var nomeusuario string
    fmt.Print("Digite seu nome: ")
    fmt.Scan(&nomeusuario)
    fmt.Println("Seja bem-vindo,", nomeusuario)

	//Scanln():
	var idadeusuario string
    fmt.Print("Digite sua idade: ")
    fmt.Scanln(&idadeusuario)
    fmt.Printf("Você tem: %s anos!", idadeusuario)

	//Scanf():
	var nomeum string
    var idadeum int

    fmt.Print("Digite seu nome e idade (ex: João 25): ")
    fmt.Scanf("%s %d", &nomeum, &idadeum)

    fmt.Printf("Nome: %s, Idade: %d\n", nomeum, idadeum)

	//ErrorF():
	err := fmt.Errorf("ocorreu um erro inesperado")
    fmt.Println(err)

	//Concatenando:
	x := "Micilini"
	fmt.Println("Olá " + "Mundo! "  + x)

	a := 10
	b := 20
	resultado := fmt.Sprint("A soma de ", a, " e ", b, " é: ", a+b)

	fmt.Println(resultado)

	nomedois := "Ana"
	idadedois := 28

	fmt.Println("Nome:", nomedois, "| Idade:", idadedois)

	nometres := "Maria"
    idadetres := 30
    salariotres := 3500.75

    // Usando fmt.Sprintf para formatar uma string
    mensagemtres := fmt.Sprintf("Olá, meu nome é %s, tenho %d anos e meu salário é R$ %.2f.", nometres, idadetres, salariotres)

    fmt.Println(mensagemtres)
}