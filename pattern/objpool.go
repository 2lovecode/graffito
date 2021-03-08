package pattern

type LinkObject struct {
	id int
}

type LinkPool chan *LinkObject

func (link *LinkObject) ID() int {
	return link.id
}

func NewPool(total int) *LinkPool {
	pool := make(LinkPool, total)

	for i := 0; i < total; i++ {
		pool <- &LinkObject{id: i}
	}

	return &pool
}

func (pool *LinkPool) Get() *LinkObject {
	select {
	case obj := <-*pool:
		*pool <- obj
		return obj
	default:
		return nil
	}
}
