package main

import (
	"fmt"
	"os"
)

func checkEnvVar(key string) {
	value, exists := os.LookupEnv(key)
	if exists {
		fmt.Printf("A variável '%s' existe e tem o valor: %s\n", key, value)
	} else {
		fmt.Printf("A variável '%s' não está definida.\n", key)
	}
}

func main() {
	os.Setenv("TOKEN_API", "abc123")

	checkEnvVar("TOKEN_API")
	checkEnvVar("NAO_EXISTE")
}