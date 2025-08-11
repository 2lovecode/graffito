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

	atomic.AddInt32(&m.waiters, 1)
	defer atomic.AddInt32(&m.waiters, -1)
	// 慢路径：进入阻塞等待
	for {
		atomic.AddInt32(&m.waiters, 1)
		if atomic.LoadInt32(&m.state) == 0 {
			if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
				return
			}
		}

		// 阻塞等待
		if err := m.sema.Acquire(context.Background(), 1); err == nil {
			// 被唤醒后尝试获取锁
			if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
				return
			}
		}
	}
}

func (m *semacas) unlock() {
	if !atomic.CompareAndSwapInt32(&m.state, 1, 0) {
		panic("unlock of unlocked mutex")
	}

	// 如果有等待者，释放信号量
	if atomic.LoadInt32(&m.waiters) > 0 {
		m.sema.Release(1)
	}
}

func newsemacas() *semacas {
	s := &semacas{
		sema:  semaphore.NewWeighted(1),
		state: 0,
	}
	// _ = s.sema.TryAcquire(1)
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
