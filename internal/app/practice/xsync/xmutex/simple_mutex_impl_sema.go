package xmutex

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"golang.org/x/sync/semaphore"
	"time"
)

type sm2 struct {
	//state int32
	sema *semaphore.Weighted
}

func new2() *sm2 {
	sema := semaphore.NewWeighted(1)
	//sema.TryAcquire(1)
	return &sm2{
		//state: 0,
		sema: sema,
	}
}

func (s *sm2) lock(ctx context.Context) error {
	//if atomic.AddInt32(&(s.state), 1) == 1 {
	//	return nil
	//}

	if err := s.sema.Acquire(ctx, 1); err != nil {
		return err
	}

	return nil
}

func (s *sm2) unlock(ctx context.Context) {
	//if atomic.AddInt32(&(s.state), -1) == 0 {
	//	return
	//}
	s.sema.Release(1)
}

type SimpleMutexImplSema struct {
}

func NewSimpleMutexImplSema() *SimpleMutexImplSema {
	return &SimpleMutexImplSema{}
}

func (s *SimpleMutexImplSema) Name() base.Name {
	return base.SimpleMutexImplSema
}

func (s *SimpleMutexImplSema) Description() string {
	return "互斥锁Mutex的简单实现：使用互斥信号量。这种实现在信号量队列头部的协程G所在的M没有获得CPU分片时间时，会阻塞后面拥有CPU分片时间的G，对CPU利用率低。"
}

func (s *SimpleMutexImplSema) Run(ctx context.Context) (string, error) {

	m2 := new2()

	c, _ := context.WithTimeout(context.Background(), 2*time.Second)

	cnt := 999
	num := 0
	total := 0

	_ = m2.lock(c)
	fmt.Println("写-已加锁")
	go func() {
		for i := 0; i <= cnt; i++ {
			num += i
		}
		fmt.Println("已写")
		m2.unlock(c)
		fmt.Println("写-已解锁")
	}()

	_ = m2.lock(c)
	fmt.Println("读-已加锁")
	total = num
	fmt.Println("已读")
	m2.unlock(c)
	fmt.Println("读-已解锁")

	return fmt.Sprintf("计算从0加到%d的总和:加锁后取结果(%d)", cnt, total), nil
}
