func main() {
    ordersA := make(chan int)
    ordersB := make(chan int)

    go feed(ordersA, 100*time.Millisecond)
    go feed(ordersB, 250*time.Millisecond)

    for i := 0; i < 5; i++ {
        select {
        case o := <-ordersA:
            fmt.Println("Pedido A:", o)
        case o := <-ordersB:
            fmt.Println("Pedido B:", o)
        case <-time.After(300 * time.Millisecond):
            fmt.Println("Nenhum pedido em 300ms — fazendo manutenção")
        }
    }
}

// feed envia inteiros a cada intervalo dado
func feed(ch chan<- int, interval time.Duration) {
    for i := 1; ; i++ {
        time.Sleep(interval)
        ch <- i
    }
}