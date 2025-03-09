package main

import "fmt"

// MapParaString recebe um map de string para string e retorna uma representação em formato de texto.
func MapParaString(m map[string]string) string {
    resultado := ""
    for chave, valor := range m {
        resultado += fmt.Sprintf("%s: %s\n", chave, valor)
    }
    return resultado
}

func main() {
    dados := map[string]string{
        "Nome":    "Carlos",
        "Idade":   "30",
        "Cidade":  "São Paulo",
    }

    resultado := MapParaString(dados)
    fmt.Println("String gerada:\n", resultado)
}