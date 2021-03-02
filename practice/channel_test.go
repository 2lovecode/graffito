package practice

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

/*
 * <-tTicket.C
 * 时钟打点channel，控制处理频率
 */
func Test_Channel_1(t *testing.T) {
	Convey("Channel_1", t, func() {
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
				fmt.Println("ch1 value: ", u)
			case v := <-ch2:
				fmt.Println("ch2 value: ", v)
			default:
				goto EE0
			}
		}
	EE0:
	})

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
func Test_Channel_2(t *testing.T) {
	Convey("Channel_2", t, func() {
		tick := time.Tick(100 * time.Millisecond)
		aTick := time.After(500 * time.Millisecond)
		for {
			select {
			case <-tick:
				fmt.Println("tick.")
			case <-aTick:
				fmt.Println("boom!")
				goto EE1
			default:
				fmt.Println(".")
				time.Sleep(500 * time.Millisecond)
			}
		}
	EE1:
	})
}

/*
 * 超时处理方案1
 */
func Test_Channel_3(t *testing.T) {
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
