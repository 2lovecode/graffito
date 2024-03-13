package help

import (
	"context"
	"github.com/2lovecode/graffito/internal/app/practice/base"
)

type QA struct {
}

func New() *QA {
	return &QA{}
}

func (qa *QA) GetQ(ctx context.Context) base.Question {
	return &Q{}
}

func (qa *QA) GetA(ctx context.Context) base.Answer {
	return &A{}
}
