package sync_x

import (
	"fmt"
	"sync"
	"time"
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

var rw sync.RWMutex
var aa int

func write() {
	rw.Lock()
	defer func() {
		rw.Unlock()
		fmt.Println("write lock end")
	}()
	fmt.Println("write lock start")
	time.Sleep(3 * time.Second)
	aa = 10
}

func read(x int) {
	rw.RLock()
	defer func() {
		rw.RUnlock()
		fmt.Printf("read lock %d end\n", x)
	}()
	fmt.Printf("read lock %d start\n", x)
	time.Sleep(1 * time.Second)
	fmt.Println("aa:", aa)

}

func Run2() {
	go write()
	time.Sleep(10 * time.Millisecond)
	go read(1)
	go read(2)

	time.Sleep(8 * time.Second)
}
