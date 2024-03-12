package qa1

type Q struct {
}

func NewQ() *Q {
	return &Q{}
}

func (q *Q) String() string {
	return "slice作为函数参数传入函数内部，并在函数内部进行操作后，原来的slice会受影响吗？"
}
