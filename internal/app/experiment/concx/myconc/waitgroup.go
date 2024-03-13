package myconc

import (
	"graffito/experiment/concx/myconc/panics"
	"sync"
)

type WaitGroup struct {
	wg sync.WaitGroup
	pc panics.Catcher
}

func (wg *WaitGroup) Go(fn func()) {
	wg.wg.Add(1)
	go func() {
		defer wg.wg.Done()
		wg.pc.Try(fn)
	}()
}

func (wg *WaitGroup) Wait() {
	wg.wg.Wait()

	wg.pc.RePanic()
}

func (wg *WaitGroup) WaitAndRecover() *panics.RecoverPanic {
	wg.wg.Wait()

	return wg.pc.Recover()
}
