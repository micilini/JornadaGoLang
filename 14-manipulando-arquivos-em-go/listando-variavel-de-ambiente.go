package main

import (
	"fmt"
	"os"
)

func listarVariaveisDeAmbiente() {
	vars := os.Environ()
	for _, v := range vars {
		fmt.Println(v)
	}
}

func main() {
	listarVariaveisDeAmbiente()
}