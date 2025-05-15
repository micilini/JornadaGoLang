package main

import (
	"fmt"
	"time"
)

func mostreAlgo(text string) <-chan bool {
	done := make(chan bool)
	go func() {
		fmt.Println(text)
		time.Sleep(5 * time.Second)
		done <- true
		close(done)
	}()
	return done
}

func main() {
	done := mostreAlgo("Micilini Roll")
	<-done
	fmt.Println("Pronto")
}