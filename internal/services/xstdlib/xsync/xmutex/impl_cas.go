package xmutex

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type cas struct {
	state int32
}

func (s *cas) lock() {
	for {
		if atomic.CompareAndSwapInt32(&s.state, 0, 1) {
			return
		}
	}
}

func (s *cas) unlock() {
	atomic.CompareAndSwapInt32(&s.state, 1, 0)
}

// 互斥锁Mutex的简单实现：用自旋CAS实现。这种实现因为在不断自旋，所以会占用大量的cpu资源
type ImplCAS struct {
}

func NewImplCAS() *ImplCAS {
	return &ImplCAS{}
}

func (s *ImplCAS) Run(ctx context.Context) {
	m1 := &cas{
		state: 0,
	}

	cnt := 0
	cnt2 := 0
	goroutines := 1000

	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	for j := 0; j < goroutines; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m1.lock()
			cnt++
			m1.unlock()
		}()
	}

	for j := 0; j < goroutines; j++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			cnt2++
		}()
	}
	wg.Wait()
	wg2.Wait()

	fmt.Printf("有锁加和{%d},无锁加和{%d}\n", cnt, cnt2)
}
