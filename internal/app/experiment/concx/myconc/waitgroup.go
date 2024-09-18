package myconc

import (
	"sync"

	"github.com/2lovecode/graffito/internal/app/experiment/concx/myconc/panics"
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
