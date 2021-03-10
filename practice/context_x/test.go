package context_x

import (
	"context"
	"time"
)

func Run1() (x1 string, x2 string){
	ctx, _ := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	ch := make(chan bool)

	go func () {
		select {
		case <- ctx.Done():
			x1 = "done"
		case <- ch:
			x1 = "ok"
		}
	}()

	go func() {
		select {
		case <- ctx.Done():
			x2 = "done"
		case <- ch:
			x2 = "ok"
		}
	}()

	<- ctx.Done()
	time.Sleep(200 * time.Millisecond)

	return x1, x2
}
