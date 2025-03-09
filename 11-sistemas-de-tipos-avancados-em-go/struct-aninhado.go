type Endereco struct {
    Rua    string
    Numero int
}

type Pessoa struct {
    Nome    string
    Idade   int
    Endereco Endereco
}

func main() {
    p := Pessoa{
        Nome:  "Lucas",
        Idade: 27,
        Endereco: Endereco{
            Rua:    "Rua das Flores",
            Numero: 123,
        },
    }

    fmt.Println(p.Nome)       // Saída: Lucas
    fmt.Println(p.Endereco.Rua) // Saída: Rua das Flores
}