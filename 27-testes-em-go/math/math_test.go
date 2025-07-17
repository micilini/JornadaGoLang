package math

import "testing"

const defaultError = "Valor Esperado %v, encontrado %v"

func TestDiscoverMedia(t *testing.T) {
    expectedValue := 7.28
    value := DiscoverMedia(7.2, 9.9, 6.1, 5.9)
    if value != expectedValue {
        t.Errorf(defaultError, expectedValue, value)
    }
}