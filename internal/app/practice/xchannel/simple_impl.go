package xchannel

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"sync"
)

type xg struct {
	f func(g *xg)
}

func newxg(f func(g *xg)) *xg {
	return &xg{f: f}
}

func (g *xg) run() {
	g.f(g)
}

type xchan struct {
	count    int
	datasize int
	sendx    int
	recvx    int
	sendq    []*xg
	recvq    []*xg
	lock     *sync.Mutex
}

func makechan() *xchan {
	return &xchan{}
}

func chansend(c *xchan, g *xg, value int) {
	// 1. 有等待接收的G，直接把值发送给队头的G
	// 2. 没有等待接收的G，有缓存则把值放入sendx指向的缓存，sendx+1
	// 3. 没有等待接收的G，没有缓存或缓存已满，则释放cpu，当前G放入sendq，并等待
}

func chanrecv(c *xchan, g *xg) (value int, err error) {
	// 1. 有等待发送的G，且无缓存，则直接把等待发送G的值copy到当前G
	// 2. 有等待发送的G，有缓存，则把等待发送G的值copy到缓存，并将recvx指向的值copy到当前G，recvx+1
	// 3. 没有等待发送的G，有缓存，则把缓存recvx指向的值copy到当前G，recvx+1
	// 4. 没有等待发送的G，没有缓存，则释放cpu，当前G放入recvq，并等待
	return
}

func closechan(c *xchan) {
	// 已经关闭的chan，报错

	// 唤醒所有的G
}

type SimpleChannelImpl struct {
}

func NewSimpleChannelImpl() *SimpleChannelImpl {
	return &SimpleChannelImpl{}
}

func (s *SimpleChannelImpl) Name() base.Name {
	return base.SimpleMutexImplCAS
}

func (s *SimpleChannelImpl) Description() string {
	return "互斥锁Mutex的简单实现：用自旋CAS实现。这种实现因为在不断自旋，所以会占用大量的cpu资源。"
}

func (s *SimpleChannelImpl) Run(ctx context.Context) (string, error) {

	gl := make([]*xg, 0, 100)
	ch := makechan()

	gl = append(gl, []*xg{
		newxg(func(g *xg) {
			v, _ := chanrecv(ch, g)
			fmt.Println(v)
		}),
		newxg(func(g *xg) {
			chansend(ch, g, 1)
		}),
	}...)

	g0 := newxg(func(g *xg) {

	})

	g0.run()

	return "", nil
}
