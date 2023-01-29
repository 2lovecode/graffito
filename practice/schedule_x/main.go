package schedule_x

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"time"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "schedule", Run: func(cmd *cobra.Command, args []string) {
		Run()
	}}
}

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
