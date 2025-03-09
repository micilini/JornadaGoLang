package main

import "fmt"

type segredo string

func (s segredo) possoEntrar(termo string) bool {
	if termo == "APTLEK" {
		return true
	}
	return false
}

func main() {
	var s segredo
	fmt.Println(s.possoEntrar("APTLEK"))
}