package modg

type Queue struct {
	len   int
	idx0  int
	idx1  int
	cap   int
	array []PkgID
}

func NewQueue(cap int) *Queue {
	o := new(Queue)

	o.cap = cap
	o.array = make([]PkgID, cap)

	return o
}

func (q *Queue) RPush(v PkgID) {
	newLen := q.len + 1
	if newLen >= q.cap {
		return
	}
	q.len = newLen
	newIdx1 := (q.idx1 + 1) % q.cap

	if newIdx1 != q.idx0 {
		q.array[newIdx1] = v
		q.idx1 = newIdx1
	}
}

func (q *Queue) LPop() PkgID {
	if q.len == 0 {
		return ""
	}

	v := q.array[q.idx0]
	q.idx0 = (q.idx0 + 1) % q.cap

	q.len = q.len - 1

	return v
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}
