package main

import (
	"fmt"
	"os"
)

func matarProcesso(pid int) {
	processo, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Processo não encontrado:", err)
		return
	}

	err = processo.Kill()
	if err != nil {
		fmt.Println("Erro ao encerrar processo:", err)
		return
	}

	fmt.Printf("Processo %d encerrado com sucesso!\n", pid)
}

func main() {
	// Altere para o PID do processo que deseja matar.
	matarProcesso(12345)
}