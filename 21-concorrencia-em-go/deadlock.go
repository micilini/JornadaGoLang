package main

import "fmt"

func main() {
    ch := make(chan int)

    ch <- 1 // Aqui o programa vai travar (deadlock)
    fmt.Println("Nunca será impresso")
}