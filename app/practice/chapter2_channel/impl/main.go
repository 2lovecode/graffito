package impl

import (
	"sync"
)

// channel实现原理

type goroutine struct {
}

type waitq struct {
	g    *goroutine
	next *waitq
}

type hchan struct {
	bufsize uint  // 缓冲区（环形队列）长度
	buf     []int // 缓冲区（环形队列），假设仅支持int
	sendx   uint  // 当前缓冲区队列下标，写入时的位置
	recvx   uint  // 当前缓冲区队列下表，读取时的位置

	sendq *waitq // 等待写入的goroutine
	recvq *waitq // 等待读取的goroutine

	lock sync.Mutex // 锁
}

func makechan(size uint) *hchan {
	var c *hchan

	c = new(hchan)
	c.bufsize = size
	c.buf = make([]int, size) // 应该是内存分配，使用slice模拟该过程

	return c
}

func read(ch *hchan) int {

	return 0
}

func write(ch *hchan, v int) {

}

func NoBuffer() {
	ch := makechan(0)
	write(ch, 1)

	read(ch)
}

func WithBuffer() {

	ch := makechan(2)

	write(ch, 1)

	read(ch)
}
