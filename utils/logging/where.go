package logging

import (
	"log"
	"runtime"
)

func Where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func F() (ret int) {
	defer func() {
		ret++
	}()

	return 1
}
