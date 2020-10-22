package main

import (
	"os"
	"runtime/trace"
)

func main() {
	// 代码是无用的
	// https://www.cnblogs.com/nickchen121/p/11517452.html
	trace.Start(os.Stderr)

	defer trace.Stop()

	ch := make(chan string)

	go func() {
		ch <- "hello world"
	}()
}
