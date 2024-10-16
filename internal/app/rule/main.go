package rule

import (
	"context"

	"github.com/2lovecode/graffito/internal/app/base"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Exec(ctx context.Context, in base.Input) (out base.Output, err error) {
	si := &Input{}
	so := &Output{}
	err = si.From(in)
	if err != nil {
		return
	}

	out = so
	return
}
