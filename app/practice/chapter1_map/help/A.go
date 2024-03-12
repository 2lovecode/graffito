package help

import (
	"context"
)

type A struct {
}

func NewA() *A {
	return &A{}
}

func (a *A) String() string {
	return "1:ab;2:ccc"
}

func (a *A) Run(ctx context.Context) string {
	return "hello worlddddd"
}
