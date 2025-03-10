package main

import (
	"fmt"
	"os"
)

func main() {
	tempDir := os.TempDir()
	fmt.Println("Diretório temporário do sistema:", tempDir)
}