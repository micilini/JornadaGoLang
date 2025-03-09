type Animal struct {
    Especie string
}

type Cachorro struct {
    Animal // Embutindo a struct Animal
    Raca   string
}

func main() {
    c := Cachorro{
        Animal: Animal{Especie: "Canino"},
        Raca:   "Labrador",
    }

    fmt.Println(c.Raca)      // Saída: Labrador
    fmt.Println(c.Especie)   // Acessa diretamente o campo da struct Animal
}