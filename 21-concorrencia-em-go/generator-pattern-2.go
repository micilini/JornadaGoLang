func countWithStop(stop <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			case <-stop:
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {
	stop := make(chan struct{})
	numbers := countWithStop(stop)

	for i := 0; i < 3; i++ {
		fmt.Println(<-numbers)
	}

	close(stop) // Interrompe o generator
}