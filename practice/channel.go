package practice

import (
	"fmt"
	"graffito/utils/params"
	"os"
	"time"
)


func ChannelRun(params params.InputParamsInterface) {
	key := params.GetInputPrefix() + "0"
	input := params.GetString(key)
	switch input {
	case "test1":
		channelTest1()
	case "test2":
		channelTest2()
	case "test3":
		channelTest3()
	default:
		fmt.Println("需要指定参数: test1/test2/test3")
	}
}
/*
 * <-tTicket.C
 * 时钟打点channel，控制处理频率
 */
func channelTest1() {
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

/*
 * time.Tick(a) 创建一个通道，每隔a时间，往通道中写入值。
 * time.After(a) 创建一个通道，在a时间后，往通道中写入值。
 */
func channelTest2() {
	tick := time.Tick(1e9)
	aTick := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-aTick:
			fmt.Println("boom!")
			os.Exit(2)
		default:
			fmt.Println(".")
			time.Sleep(5e8)
		}
	}
}


/*
 * 超时处理方案1
 */
func channelTest3() {
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

