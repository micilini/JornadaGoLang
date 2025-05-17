func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	for result := range square(gen()) {
		fmt.Println(result)
	}
}