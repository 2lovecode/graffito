package main

import (
	"fmt"
	"time"
)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go push(ch1)
	go push(ch2)
	go print(ch1, ch2)
	time.Sleep(3e5)
}

func push(ch chan int) {
	for i := 0; i < 10; i ++ {
		ch <- i
	}
}

func print(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1 value :", v)
		case v := <-ch2:
			fmt.Println("ch2 value :", v)
		default:
			fmt.Println("waiting")
		}
	}
}
