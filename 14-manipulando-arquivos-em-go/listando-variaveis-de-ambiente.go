package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Variáveis de ambiente:")

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}