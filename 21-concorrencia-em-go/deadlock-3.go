func main() {
    a := make(chan int)
    b := make(chan int)

    go func() {
        <-b // espera receber algo de b
        a <- 1
    }()

    <-a // espera receber algo de a
}