package panics

import "sync/atomic"

type Catcher struct {
	recovered atomic.Pointer[any]
}

func (cat *Catcher) Try(fn func()) {

}

func (cat *Catcher) RePanic() {

}

func (cat *Catcher) Recover() *RecoverPanic {
	return nil
}

type RecoverPanic struct {
}

func (rp *RecoverPanic) Error() string {
	return ""
}
