package sync_x

import (
	"fmt"
	"sync"
)

func Run1() {
	var once sync.Once
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Hello World!")
			})
			ch <- true
		}()
	}

	go func() {
		close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

}
