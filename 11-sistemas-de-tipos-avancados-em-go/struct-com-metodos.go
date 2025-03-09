type Pessoa struct {
    Nome string
}

// Método associado à struct Pessoa
func (p Pessoa) Saudacao() {
    fmt.Println("Olá,", p.Nome)
}

func main() {
    p := Pessoa{"Amanda"}
    p.Saudacao() // Saída: Olá, Amanda
}