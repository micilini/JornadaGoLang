package main

import (
    "fmt"
    "time"
)

func main() {
	now := time.Now()

	// 1. Em segundos (Unix epoch, 10 dígitos aprox.)
	tsSegundos := now.Unix()                
	fmt.Println("Timestamp (segundos):", tsSegundos)

	// 2. Em nanosegundos (Unix epoch, 19 dígitos aprox.)
	tsNano := now.UnixNano()                 
	fmt.Println("Timestamp (nanosegundos):", tsNano)

	// 3. Em milissegundos (Unix epoch, 13 dígitos)
	// Basta dividir o UnixNano por 1_000_000
	tsMillis := now.UnixNano() / int64(time.Millisecond)
	fmt.Println("Timestamp (milissegundos):", tsMillis)
}