package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for i := range ch {
		fmt.Println("Reseive:", i)
	}
}

func main() {
	ch := make(chan int, 10)
	go produce(ch)
	consume(ch)
	time.Sleep(1 * time.Second)
}
