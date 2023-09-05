package controllers

import (
	"context"
	"github.com/kataras/iris/v12"
	"graffito/app/sandbox"
	"graffito/pkg/response"
)

type Sandbox struct{}

func NewSandboxController() *Sandbox {
	return &Sandbox{}
}

type RunJson struct {
	SourceCode string `json:"sourceCode"`
}

type RunData struct {
	Content string `json:"content"`
}

func (sd *Sandbox) Run(ctx iris.Context) {
	runJ := RunJson{}
	err := ctx.ReadJSON(&runJ)
	if err != nil {
		response.Failure(ctx, "400", err.Error())
		return
	}

	sandboxApp := sandbox.NewApplication()
	out, err := sandboxApp.Exec(context.Background(), &sandbox.Input{
		SourceCode: runJ.SourceCode,
	})

	if err != nil {
		response.Failure(ctx, "400", err.Error())
		return
	}
	so := sandbox.Output{}
	err = out.To(&so)

	if err != nil {
		response.Failure(ctx, "400", err.Error())
		return
	}

	response.Success(ctx, &RunData{
		Content: so.Data,
	})
}
