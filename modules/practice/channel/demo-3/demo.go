package main

import (
	"fmt"
	"time"
)

/*
 * 超时处理方案1
 */
func main() {
	timeout := make(chan int)
	defer close(timeout)

	go func() {
		//超时时间设置为1s
		time.Sleep(1e9)
		timeout <- 1
	}()

	ch := make(chan int)
	defer close(ch)

	go func() {
		time.Sleep(2e9)
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("right!")
	case <-timeout:
		fmt.Println("timeout!")
	}
}
