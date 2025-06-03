package main

import (
    "fmt"
    "time"
)

func main() {
	t1 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now().UTC()

	if t1.Before(t2) {
		fmt.Println("t1 é anterior a t2")
	} else if t1.After(t2) {
		fmt.Println("t1 é posterior a t2")
	} else {
		fmt.Println("t1 e t2 são iguais")
	}
}