package main

import (
	"fmt"
	"os"
)

func mostrarPID() {
	fmt.Printf("PID (Processo Atual): %d\n", os.Getpid())
	fmt.Printf("PPID (Processo Pai): %d\n", os.Getppid())
}

func main() {
	mostrarPID()
}