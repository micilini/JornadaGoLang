package main

import (
    "fmt"
    "time"
)

func main() {
	// Exemplo: converter timestamp Unix para Time
	ts := time.Now().Unix()
	tFromUnix := time.Unix(ts, 0) // 0 nanosegundos extra
	fmt.Println("De volta:", tFromUnix)

	// Mesmo com nanosegundos
	tsNano := time.Now().UnixNano()
	tFromNano := time.Unix(0, tsNano)
	fmt.Println("De volta (nano):", tFromNano)
}