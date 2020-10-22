package lib

import "fmt"

type LinkObject struct {
	id int
}


type LinkPool chan *LinkObject

func (link *LinkObject) ID() {
	fmt.Println("ID is :", link.id)
}

func NewPool(total int) *LinkPool {
	pool := make(LinkPool, total)

	for i := 0; i < total; i++ {
		pool <- &LinkObject{id: i}
	}

	return &pool
}


func (pool *LinkPool) Do() {
	select {
	case obj := <-*pool:
		obj.ID()
		*pool <- obj
	default:
		return
	}
}

