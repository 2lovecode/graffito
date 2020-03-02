package practice

import (
	"fmt"
	"graffito/utils/params"
	"time"
)

func RoutineRun(p params.InputParamsInterface) {
	key := p.GetInputPrefix() + "0"
	input := p.GetString(key)
	switch input {
	case "test1":
		routineTest1()
	case "test2":
		routineTest2()
	default:
		fmt.Println("请输入参数 test1/test2")
	}
}

func routineTest1() {
	out := routineSieve()
	for v := range out{
		fmt.Print(v, " ")
	}
}

func routineCreate() (in chan int) {
	in = make(chan int)
	go func() {
		for i := 2; i < 6;i++ {
			in <- i
		}
		close(in)
	}()
	return
}

func routineFilter(in chan int, c int) (out chan int){
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

func routineSieve() (out chan int){
	out = make(chan int)
	go func() {
		in := routineCreate()
		for {
			v := <- in
			in = routineFilter(in, v)
			out <- v
		}
	}()
	return
}


func routineTest2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go routinePush(ch1)
	go routinePush(ch2)
	go routinePrint(ch1, ch2)
	time.Sleep(3e5)
}

func routinePush(ch chan int) {
	for i := 0; i < 10; i ++ {
		ch <- i
	}
}

func routinePrint(ch1 chan int, ch2 chan int) {
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
