package link

import "errors"

type StackErrFull error
type StackErrEmpty error

type Stack[T any] struct {
	cap uint
	len uint
	arr []T
}

func NewStack[T any](cap uint) *Stack[T] {
	o := new(Stack[T])
	o.arr = make([]T, cap)
	o.cap = cap
	o.len = 0
	return o
}

func (s *Stack[T]) Empty() bool {
	return s.len <= 0
}

func (s *Stack[T]) Full() bool {
	return s.cap == s.len
}

func (s *Stack[T]) Push(v T) error {
	if s.Full() {
		return StackErrFull(errors.New("栈已满"))
	}
	s.arr[s.len] = v
	s.len++
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var v T
	if s.Empty() {
		return v, StackErrEmpty(errors.New("栈为空"))
	}
	s.len--
	v = s.arr[s.len]
	return v, nil
}
