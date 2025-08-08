package sandbox

import (
	"context"
	"os"

	dto "github.com/2lovecode/graffito/internal/dto/sandbox"
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

func (sand *Application) Exec(ctx context.Context, in dto.Input) (out dto.Output, err error) {
	outIO := &StringIO{data: make([]byte, 0)}

	inter := interp.New(interp.Options{
		GoPath: os.Getenv("GOPATH"),
		Stdout: outIO,
	})

	err = inter.Use(stdlib.Symbols)

	if err != nil {
		return
	}

	_, err = inter.Eval(in.SourceCode)

	if err != nil {
		return
	}

	out = dto.Output{
		Data: outIO.String(),
	}
	return
}

func NewApplication() *Application {
	return &Application{}
}
