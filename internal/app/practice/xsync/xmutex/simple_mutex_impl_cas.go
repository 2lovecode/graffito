package xmutex

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"sync/atomic"
)

type sm1 struct {
	state int32
}

func (s *sm1) lock() {
	for {
		fmt.Println("before-lock", s.state)
		if atomic.CompareAndSwapInt32(&s.state, 0, 1) {
			fmt.Println("after-lock", s.state)
			return
		}
	}
}

func (s *sm1) unlock() {
	fmt.Println("before-unlock", s.state)
	atomic.CompareAndSwapInt32(&s.state, 1, 0)
	fmt.Println("after-unlock", s.state)
}

type SimpleMutexImplCAS struct {
}

func NewSimpleMutexImpl() *SimpleMutexImplCAS {
	return &SimpleMutexImplCAS{}
}

func (s *SimpleMutexImplCAS) Name() base.Name {
	return base.SimpleMutexImplCAS
}

func (s *SimpleMutexImplCAS) Description() string {
	return "互斥锁Mutex的简单实现：用自旋CAS实现。这种实现因为在不断自旋，所以会占用大量的cpu资源。"
}

func (s *SimpleMutexImplCAS) Run(ctx context.Context) (string, error) {
	m1 := &sm1{}

	cnt := 999
	num := 0
	total := 0

	m1.lock()
	fmt.Println("写-已加锁")
	go func() {
		for i := 0; i <= cnt; i++ {
			num += i
		}
		fmt.Println("已写")
		m1.unlock()
		fmt.Println("写-已解锁")
	}()

	m1.lock()
	fmt.Println("读-已加锁")
	total = num
	fmt.Println("已读")
	m1.unlock()
	fmt.Println("读-已解锁")

	return fmt.Sprintf("计算从0加到%d的总和:加锁后取结果(%d)", cnt, total), nil
}
