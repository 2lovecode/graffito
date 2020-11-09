package main

import (
	"graffito/pattern/lib"
	"time"
)

func main() {
	notifier := lib.NewNotifier()

	notifier.Register(lib.NewObserver(1))
	notifier.Register(lib.NewObserver(2))

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <- stop:
			return
			case t := <-tick:
				notifier.Notify(lib.Event{Data:t.UnixNano()})
		}
	}
}
