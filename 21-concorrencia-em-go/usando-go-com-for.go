package strings

import (
    "testing"
    "strings"
)

// msgIndex é o template de erro com placeholders:
// %s → texto completo
// %s → substring buscada
// %d → índice esperado
// %d → índice encontrado
const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
    // 1) Definição do “dataset”: slice de structs com três campos
    tests := []struct {
        text     string // texto onde vamos buscar
        part     string // substring a buscar
        expected int    // índice esperado de retorno
    }{
        {"Micilini is Best!", "Micilini", 0},  // começa na posição 0
        {"", "", 0},                           // string vazia em string vazia → índice 0
        {"Hey", "hey", -1},                    // busca case‑sensitive: não encontra → -1
        {"GoLang", "L", 2},                    // “L” começa no índice 2 em “GoLang”
    }

    // 2) Loop que executa cada caso de teste
    for _, test := range tests {
        // Loga o caso atual (útil para debugging)
        t.Logf("Result: %v", test)

        // 3) Chama a função que está sob teste
        current := strings.Index(test.text, test.part)

        // 4) Compara resultado com o esperado
        if current != test.expected {
            // Emite um erro formatado se der mismatch
            t.Errorf(
                msgIndex,
                test.text,
                test.part,
                test.expected,
                current,
            )
        }
    }
}