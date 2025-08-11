package xmutex

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type semacas struct {
	state   int32
	mu      sync.Mutex // 仅示例
	waiters int32
	sema    *semaphore.Weighted
}

func (s *semacas) print() {
	fmt.Println("state:", atomic.LoadInt32(&s.state))
}
func (m *semacas) lock() {
	// 快速路径：尝试CAS获取锁
	if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
		return
	}

	// 自旋尝试
	for i := 0; i < 10; i++ {
		if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			return
		}
	}

	// 慢路径：进入阻塞等待
	for {
		m.mu.Lock()
		m.waiters++
		m.mu.Unlock()
		if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			m.mu.Lock()
			m.waiters--
			m.mu.Unlock()
			return
		}

		// 阻塞等待
		if err := m.sema.Acquire(context.Background(), 1); err != nil {
			panic(err.Error())
		}
		if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			return
		}
	}
}

func (m *semacas) unlock() {
	if !atomic.CompareAndSwapInt32(&m.state, 1, 0) {
		panic("unlock of unlocked mutex")
	}

	m.mu.Lock()
	if m.waiters > 0 {
		m.waiters--
		m.mu.Unlock()
		m.sema.Release(1)
	} else {
		m.mu.Unlock()
	}
}

func newsemacas() *semacas {
	s := &semacas{
		sema:  semaphore.NewWeighted(1),
		state: 0,
	}
	if !s.sema.TryAcquire(1) {
		panic("init semaphore fail")
	}
	return s
}

// 互斥锁Mutex的简单实现：使用CAS+互斥信号量实现。这种实现给到新协程竞争锁的机会，提升了cpu利用率。但是在高并发情况下会存在惊群现象（大量被唤醒的协程在同时竞争锁）。
type ImplSemaCAS struct {
}

func NewImplSemaCAS() *ImplSemaCAS {
	return &ImplSemaCAS{}
}

func (s *ImplSemaCAS) Run(ctx context.Context) {
	m1 := newsemacas()

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
			// m1.print()
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
