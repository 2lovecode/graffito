package qa1

import (
	"context"
)

type A struct {
}

func NewA() *A {
	return &A{}
}

func (a *A) String() string {
	return "aa"
}

func (a *A) Run(ctx context.Context) string {
	return "hello worlddddd"
}
