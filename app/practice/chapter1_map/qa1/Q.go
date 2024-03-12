package qa1

type Q struct {
}

func NewQ() *Q {
	return &Q{}
}

func (q *Q) String() string {
	return "a"
}
