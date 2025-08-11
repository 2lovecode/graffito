package xmutex

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	Locked      int32 = 1
	Woken       int32 = 2
	WaiterShift int32 = 2
)

type sm4 struct {
	state int32
	sema  *semaphore.Weighted
}

func new4() *sm4 {
	sema := semaphore.NewWeighted(1)
	sema.TryAcquire(1)
	return &sm4{
		state: Locked,
		sema:  sema,
	}
}

func (s *sm4) lock(ctx context.Context) error {
	if atomic.CompareAndSwapInt32(&(s.state), 0, Locked) {
		return nil
	}
	woken := false
	for {
		ov := s.state

		nv := ov | Locked // 加锁

		if ov&Locked != 0 {
			nv = ov + 1<<WaiterShift
		}

		if woken {
			nv &^= Woken
		}

		if atomic.CompareAndSwapInt32(&(s.state), ov, nv) {
			if ov&Locked == 0 {
				break
			}

			if err := s.sema.Acquire(ctx, 1); err != nil {
				return err
			}

			woken = true
		}
	}

	return nil
}

func (s *sm4) unlock(ctx context.Context) {

	nv := atomic.AddInt32(&(s.state), -Locked)

	ov := nv

	for {
		if ov>>WaiterShift == 0 || (ov&(Locked|Woken) != 0) {
			return
		}

		nv = (ov - 1<<WaiterShift) | Woken

		if atomic.CompareAndSwapInt32(&(s.state), ov, nv) {
			s.sema.Release(1)
			return
		}
		ov = s.state
	}
}

// "基于 CAS2 的信号量实现"
type SimpleMutexImplSemaCas2 struct{}

func NewSimpleMutexImplSemaCas2() *SimpleMutexImplSemaCas2 {
	return &SimpleMutexImplSemaCas2{}
}

func (s *SimpleMutexImplSemaCas2) Run(ctx context.Context) (string, error) {
	m4 := new4()

	c, _ := context.WithTimeout(context.Background(), 2*time.Second)

	cnt := 999
	num := 0
	total := 0

	_ = m4.lock(c)
	fmt.Println("写-已加锁")
	go func() {
		for i := 0; i <= cnt; i++ {
			num += i
		}
		fmt.Println("已写")
		m4.unlock(c)
		fmt.Println("写-已解锁")
	}()

	_ = m4.lock(c)
	fmt.Println("读-已加锁")
	total = num
	fmt.Println("已读")
	m4.unlock(c)
	fmt.Println("读-已解锁")

	return fmt.Sprintf("计算从0加到%d的总和:加锁后取结果(%d)", cnt, total), nil
}
