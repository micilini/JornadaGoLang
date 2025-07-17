package math

import (
    "fmt"
    "strconv"
)

// Calcula a média entre dois ou mais valores recebidos por parâmetros
func DiscoverMedia(numbers ...float64) float64 {
    total := 0.0
    for _, num := range numbers {
        total += num
    }
    media := total / float64(len(numbers))
    // formata com duas casas decimais
    result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
    return result
}