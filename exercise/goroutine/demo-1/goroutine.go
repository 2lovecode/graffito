package main

import (
	"fmt"
)
func main() {
	out := sieve()
	for v := range out{
		fmt.Print(v, " ")
	}
}

func create() (in chan int) {
	in = make(chan int)
	go func() {
		for i := 2; i < 6;i++ {
			in <- i
		}
		close(in)
	}()
	return
}

func filter(in chan int, c int) (out chan int){
	out = make(chan int)
	go func() {
		for v := range in{
			if v%c != 0 {
				out <- v
			}
		}
	}()
	return
}

func sieve() (out chan int){
	out = make(chan int)
	go func() {
		in := create()
		for {
			v := <- in
			in = filter(in, v)
			out <- v
		}
	}()
	return
}