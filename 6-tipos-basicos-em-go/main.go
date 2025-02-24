package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Exemplos de string
	str1 := "Olá, Mundo!"
	str2 := "Micilini"
	strMultiline := `Olá
, Portal
da Micilini!`
	fmt.Println(str1, str2, strMultiline)

	// Exemplos de números inteiros
	intEx := 42
	int8Ex := int8(-128)
	int16Ex := int16(32000)
	int32Ex := int32(100000)
	int64Ex := int64(9223372036854775807)
	fmt.Println(intEx, int8Ex, int16Ex, int32Ex, int64Ex)

	// Inteiros sem sinal
	uintEx := uint(42)
	uint8Ex := uint8(255)
	uint16Ex := uint16(65535)
	uint32Ex := uint32(400000)
	uint64Ex := uint64(18446744073709551615)
	fmt.Println(uintEx, uint8Ex, uint16Ex, uint32Ex, uint64Ex)

	// byte e rune
	byteEx := byte(65)
	runeEx := rune('G')
	fmt.Println(byteEx, runeEx)

	// Números de ponto flutuante
	float32Ex := float32(3.1415927)
	float64Ex := 3.141592653589793
	fmt.Println(float32Ex, float64Ex)

	// Números complexos
	complex64Ex := complex(float32(3.5), float32(4.2))
	complex128Ex := complex(5.5, 6.7)
	fmt.Println(complex64Ex, complex128Ex)

	// Booleanos
	boolTrue := true
	boolFalse := false
	fmt.Println(boolTrue, boolFalse)

	// Usando reflect
	fmt.Println(reflect.TypeOf(intEx))
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.TypeOf(boolTrue))

	// Tipos zero
	var zeroString string
	var zeroInt int
	var zeroFloat64 float64
	var zeroBool bool
	fmt.Println(zeroString, zeroInt, zeroFloat64, zeroBool)

	// Conversões entre tipos
	intToFloat := float64(intEx)
	floatToInt := int(float64Ex)
	byteToString := string([]byte{71, 111})
	fmt.Println(intToFloat, floatToInt, byteToString)

	// Usando strconv para converter string para int
	strNum := "42"
	intFromStr, err := strconv.Atoi(strNum)
	if err == nil {
		fmt.Println("Convertido com sucesso:", intFromStr)
	}

	// Convertendo int para string
	strFromInt := strconv.Itoa(intEx)
	fmt.Println("String do int:", strFromInt)

	// Convertendo string para bool
	boolStr, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("Booleano convertido:", boolStr)
	}
}