package chapter2_channel

import (
	"time"
)

/*
 * <-tTicket.C
 * 时钟打点channel，控制处理频率
 */
func Run1() (x1, x2 int) {
	tTicket := time.NewTicker(1000 * time.Millisecond)
	defer tTicket.Stop()
	ch1 := make(chan int)
	ch2 := make(chan int)

	go go1(ch1)
	go go2(ch2)

	for {
		<-tTicket.C
		select {
		case u := <-ch1:
			x1 = u
		case v := <-ch2:
			x2 = v
		default:
			goto EE0
		}
	}
EE0:
	return
}

func go1(ch chan int) {
	ch <- 1
}

func go2(ch chan int) {
	ch <- 2
}

/*
 * time.Tick(a) 创建一个通道，每隔a时间，往通道中写入值。
 * time.After(a) 创建一个通道，在a时间后，往通道中写入值。
 */
func Run2() (x1, x2 string) {
	tick := time.Tick(10 * time.Millisecond)
	aTick := time.After(50 * time.Millisecond)
	for {
		select {
		case <-tick:
			x1 = "tick"
		case <-aTick:
			x2 = "boom"
			goto EE1
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
EE1:
	return
}

/*
 * 超时处理方案1
 */
func Run3() (r string) {
	timeout := make(chan int)

	go func() {
		//超时时间设置为10ms
		time.Sleep(10 * time.Millisecond)
		timeout <- 1
	}()

	ch := make(chan int)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch <- 1
	}()

	select {
	case <-ch:
		r = "right"
	case <-timeout:
		r = "timeout"
	}
	return
}
