package main

import (
	"context"
	"sync"
	"time"
)

func main() {

	timeOutTime := time.Now().Add(500 * time.Millisecond)
	timeOut, _ := context.WithDeadline(context.Background(), timeOutTime)

	var wg sync.WaitGroup


	pics := []string{"pic1", "pic2"}

	for _, eachPic := range pics {
		wg.Add(1)
		go func (pic string) {
			ch := make(chan bool)
			go DownloadPic(pic, ch)
			for {
				select {
				case <-ch:
					wg.Done()
					return
				case <-timeOut.Done():
					wg.Done()
					return
				}
			}
		}(eachPic)
	}

	wg.Wait()

}

func DownloadPic(pic string, ch chan bool) {
	// http.Get(pic)
	ch <- true
}




