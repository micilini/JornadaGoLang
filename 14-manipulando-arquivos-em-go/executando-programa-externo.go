package main

import (
	"fmt"
	"os"
)

func abrirBlocoDeNotas() {
	programa := "notepad.exe"
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	processo, err := os.StartProcess(programa, []string{programa}, attr)
	if err != nil {
		fmt.Println("Erro ao iniciar processo:", err)
		return
	}

	fmt.Printf("Bloco de Notas iniciado com PID: %d\n", processo.Pid)
}

func main() {
	abrirBlocoDeNotas()
}