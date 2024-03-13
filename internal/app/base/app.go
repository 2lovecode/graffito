package base

import "context"

type Application interface {
	Exec(ctx context.Context, in Input) (out Output, err error)
}
