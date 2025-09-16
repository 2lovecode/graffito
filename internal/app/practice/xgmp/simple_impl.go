package xgmp

import (
	"context"

	"github.com/2lovecode/graffito/internal/app/practice/base"
)

var queue = make([]*xg, 100)

type xm struct {
}

type AA struct {
	M string
	S string
}
type xp struct {
	m     *xm
	queue []*xg
}

type xg struct {
}

// schedule 调度器
func schedule() {

}

func syscall(f func()) {
	entersyscall()
	f()
	exitsyscall()
}

func entersyscall() {
	//
}

func exitsyscall() {
	// 触发重新调度
	schedule()
}

type SimpleGMPlImpl struct {
}

func NewSimpleGMPlImpl() *SimpleGMPlImpl {
	return &SimpleGMPlImpl{}
}

func (s *SimpleGMPlImpl) Name() base.Name {
	return base.SimpleGMPImpl
}

func (s *SimpleGMPlImpl) Description() string {
	return "互斥锁Mutex的简单实现：用自旋CAS实现。这种实现因为在不断自旋，所以会占用大量的cpu资源。"
}

func (s *SimpleGMPlImpl) Run(ctx context.Context) (string, error) {

	// 用Goroutine来模拟系统线程 M

	return "", nil
}
