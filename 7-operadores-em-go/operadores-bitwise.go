package main

import "fmt"

func main() {
    a := 5  // 101 em binário
    b := 3  // 011 em binário

    fmt.Println("a AND b:", a & b)      // 1 (0001)
    fmt.Println("a OR b:", a | b)       // 7 (0111)
    fmt.Println("a XOR b:", a ^ b)      // 6 (0110)
    fmt.Println("a AND NOT b:", a &^ b) // 4 (0100)

    fmt.Println("a << 1 (Left Shift):", a << 1) // 10 (1010)
    fmt.Println("a >> 1 (Right Shift):", a >> 1) // 2 (0010)
}