package sandbox

import (
	"context"
	"github.com/2lovecode/graffito/internal/app/base"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type Application struct {
}

type StringIO struct {
	data []byte
}

func (so *StringIO) Write(p []byte) (n int, err error) {
	so.data = append(so.data, p...)
	return len(p), nil
}

func (so *StringIO) Read(p []byte) (n int, err error) {
	p = append(p, so.data...)
	return len(p), nil
}

func (so *StringIO) String() string {
	return string(so.data)
}

func (sand *Application) Exec(ctx context.Context, in base.Input) (out base.Output, err error) {
	si := &Input{}
	so := &Output{}
	err = si.From(in)

	if err != nil {
		return
	}

	outIO := &StringIO{data: make([]byte, 0)}

	inter := interp.New(interp.Options{
		Stdout: outIO,
	})

	err = inter.Use(stdlib.Symbols)

	if err != nil {
		return
	}

	_, err = inter.Eval(si.SourceCode)

	if err != nil {
		return
	}
	so.Data = outIO.String()

	out = so
	return
}

func NewApplication() *Application {
	return &Application{}
}
