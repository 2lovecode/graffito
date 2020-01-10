package main

import (
	"fmt"
	"os"
	"time"
)

/*
 * time.Tick(a) 创建一个通道，每隔a时间，往通道中写入值。
 * time.After(a) 创建一个通道，在a时间后，往通道中写入值。
 */
func main() {
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
