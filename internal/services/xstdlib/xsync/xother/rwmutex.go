package xother

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

const rwmutexMaxReaders = 1 << 30

type RWMutex struct {
	mu                sync.Mutex
	readerSem         semaphore.Weighted
	readerSemWaiterMu sync.Mutex
	readerSemWaiter   int
	writerSem         semaphore.Weighted
	writerSemWaiterMu sync.Mutex
	writerSemWaiter   int
	readerCount       int32
	writerCount       int32
}

func NewRWMutex() *RWMutex {
	return &RWMutex{}
}

func (rw *RWMutex) Lock() {
	rw.mu.Lock()
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	if r != 0 {
		rw.writerSemWaiterMu.Lock()
		rw.writerSemWaiter++
		rw.writerSemWaiterMu.Unlock()
		rw.writerSem.Acquire(context.Background(), 1)
	}
}
func (rw *RWMutex) Unlock() {
	atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)

	rw.readerSemWaiterMu.Lock()
	if rw.readerSemWaiter > 0 {
		rw.readerSemWaiter--
		rw.readerSem.Release(1)
	}
	rw.readerSemWaiterMu.Unlock()
	rw.mu.Unlock()
}
func (rw *RWMutex) RLock() {
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// 等待写锁释放
		rw.readerSemWaiterMu.Lock()
		rw.readerSemWaiter++
		rw.readerSemWaiterMu.Unlock()

		rw.readerSem.Acquire(context.Background(), 1)
	}
}
func (rw *RWMutex) RUnlock() {
	if atomic.AddInt32(&rw.readerCount, -1) < 0 {
		rw.writerSemWaiterMu.Lock()
		if rw.writerSemWaiter > 0 {
			rw.writerSemWaiter--
			rw.writerSem.Release(1)
		}
		rw.writerSemWaiterMu.Unlock()
	}
}
