package xmutex

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/semaphore"
)

// 关于信号量
// Go的信号量实现是在用户态实现的轻量级信号量，它的上下文切换相比OS的信号量较轻，性能更好。

// 基础阻塞锁
// 使用信号量实现
// 1. 性能
// a.信号量队列头部的G，若它所在的M没有获取到CPU分片，则会阻塞。此时若队列里的其他G所在的M获取到了CPU分片时间，它也不会被运行，在多核CPU场景下，这种情况更普遍。
// b.锁竞争时，立即进入阻塞。在短等待的锁场景下，不够高效。
// 2. 锁公平
// 新的G来竞争锁可能会插队（G的调度问题）

type sema struct {
	sm *semaphore.Weighted // 仅用作示例
}

func (s *sema) lock() {
	s.sm.Acquire(context.Background(), 1)
}

func (s *sema) unlock() {
	s.sm.Release(1)
}

type ImplSema struct {
}

func NewImplSema() *ImplSema {
	return &ImplSema{}
}

func (s *ImplSema) Run(ctx context.Context) {
	m1 := &sema{
		sm: semaphore.NewWeighted(1),
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
