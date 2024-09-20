package xchannel

import (
	"fmt"
	"sync"
)

// goroutine + channel 实现并发的几个模型
// 1. 1-1
// 2. 1-N
// 3. N-1
// 4. N-N

func OneToOne() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	go func() {
		a := <-ch
		fmt.Println(a)
	}()
}

func OneToMany() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for j := 0; j < 2; j++ {
		go func() {
			a := <-ch
			fmt.Println(a)
		}()
	}
}

func ManyToOne() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			ch <- i
		}(i)
	}

	go func() {
		wg.Done()
		close(ch)
	}()

	go func() {
		for a := range ch {
			fmt.Println(a)
		}
	}()
}

func ManyToMany() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			ch <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for j := 0; j < 2; j++ {
		go func() {
			for a := range ch {
				fmt.Println(a)
			}
		}()
	}
}
