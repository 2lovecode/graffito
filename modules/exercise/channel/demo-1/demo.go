package main


import (
	"fmt"
	"os"
	"time"
)

/*
 * <-tTicket.C
 * 时钟打点channel，控制处理频率
 */
func main() {
	tTicket := time.NewTicker(1e9)
	defer tTicket.Stop()
	ch1 := make(chan int)
	ch2 := make(chan int)

	go go1(ch1)
	go go2(ch2)

	for {
		<-tTicket.C
		select {
		case u := <- ch1:
			fmt.Println("ch1 value: ", u)
		case v := <- ch2:
			fmt.Println("ch2 value: ", v)
		default:
			os.Exit(1)
		}
	}
}

func go1(ch chan int) {
	ch <- 1
}

func go2(ch chan int) {
	ch <- 2
}
