package xmutex

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"golang.org/x/sync/semaphore"
	"sync/atomic"
	"time"
)

type sm3 struct {
	state int32
	sema  *semaphore.Weighted
}

func new3() *sm3 {
	sema := semaphore.NewWeighted(1)
	sema.TryAcquire(1)

	return &sm3{
		state: 0,
		sema:  sema,
	}
}

func (s *sm3) lock(ctx context.Context) error {
	if atomic.CompareAndSwapInt32(&(s.state), 0, 1) {
		return nil
	}

	for {
		if atomic.CompareAndSwapInt32(&(s.state), 0, 1) {
			return nil
		}
		if err := s.sema.Acquire(ctx, 1); err != nil {
			return err
		}
	}
}

func (s *sm3) unlock(ctx context.Context) {
	if atomic.CompareAndSwapInt32(&(s.state), 1, 0) {
		s.sema.Release(1)
		return
	}
}

type SimpleMutexImplSemaCAS struct {
}

func NewSimpleMutexImplSemaCAS() *SimpleMutexImplSemaCAS {
	return &SimpleMutexImplSemaCAS{}
}

func (s *SimpleMutexImplSemaCAS) Name() base.Name {
	return base.SimpleMutexImplSemaCAS
}

func (s *SimpleMutexImplSemaCAS) Description() string {
	return "互斥锁Mutex的简单实现：使用CAS+互斥信号量实现。这种实现给到新协程竞争锁的机会，提升了cpu利用率。但是在高并发情况下会存在惊群现象（大量被唤醒的协程在同时竞争锁）。"
}

func (s *SimpleMutexImplSemaCAS) Run(ctx context.Context) (string, error) {
	m3 := new3()

	c, _ := context.WithTimeout(context.Background(), 2*time.Second)

	cnt := 999
	num := 0
	total := 0

	_ = m3.lock(c)
	fmt.Println("写-已加锁")
	go func() {
		for i := 0; i <= cnt; i++ {
			num += i
		}
		fmt.Println("已写")
		m3.unlock(c)
		fmt.Println("写-已解锁")
	}()

	_ = m3.lock(c)
	fmt.Println("读-已加锁")
	total = num
	fmt.Println("已读")
	m3.unlock(c)
	fmt.Println("读-已解锁")

	return fmt.Sprintf("计算从0加到%d的总和:加锁后取结果(%d)", cnt, total), nil
}
