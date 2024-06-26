package practice

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/base"
	base2 "github.com/2lovecode/graffito/internal/app/practice/base"
	"github.com/2lovecode/graffito/internal/app/practice/xsync/xmutex"
)

type Application struct {
	all map[base2.Name]base2.Question
}

func NewApplication() *Application {
	app := &Application{}
	app.init()

	return app
}

func (app *Application) Exec(ctx context.Context, in base.Input) (out base.Output, err error) {
	si := &Input{}
	so := &Output{}
	err = si.From(in)
	if err != nil {
		return
	}

	if si.List {
		for _, v := range app.all {
			so.Data += fmt.Sprintf("\nQ: %s\nD: %s\n", v.Name(), v.Description())
		}
	} else if c, ok := app.all[base2.Name(si.Question)]; ok {
		so.Data, so.Error = c.Run(ctx)
	}
	out = so
	return
}

func (app *Application) init() {
	xmutex1 := xmutex.NewSimpleMutexImpl()
	app.all = map[base2.Name]base2.Question{
		xmutex1.Name(): xmutex1,
	}
}
