package schedule_x

import (
	"fmt"
	"runtime"
	"time"
)

func Run() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("A: ", i)
		}()
	}

	// var ch = make(chan int)
	// <-ch

	time.Sleep(time.Hour)
}
