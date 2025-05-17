package main

import (
	"fmt"
	"time"
)

func producer(name string, delay time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(delay)
			ch <- fmt.Sprintf("%s produz: %d", name, i)
		}
		close(ch)
	}()
	return ch
}

func multiplex(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for ch1 != nil || ch2 != nil {
			select {
			case val, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- val
			case val, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- val
			}
		}
	}()
	return out
}

func main() {
	a := producer("A", 500*time.Millisecond)
	b := producer("B", 700*time.Millisecond)

	mux := multiplex(a, b)

	for val := range mux {
		fmt.Println(val)
	}
}