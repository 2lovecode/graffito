package practice

import (
	"context"
	"github.com/2lovecode/graffito/internal/app/base"
	base2 "github.com/2lovecode/graffito/internal/app/practice/base"
	"github.com/2lovecode/graffito/internal/app/practice/chapter1_map"
	"github.com/2lovecode/graffito/internal/app/practice/chapter2_slice"
	"github.com/2lovecode/graffito/pkg/json"
)

type Application struct {
	chapter map[string]base2.Chapter
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

	if c, ok := app.chapter[si.Chapter]; ok {
		qa := c.GetQA(ctx, si.Number)
		if qa != nil {
			switch si.Part {
			case "q":
				q := qa.GetQ(ctx)
				if q != nil {
					so.Data = q.String()
				}
			case "a":
				a := qa.GetA(ctx)
				if a != nil {
					so.Data = a.String()
				}
			case "all":
				all := make(map[string]string)
				q := qa.GetQ(ctx)
				if q != nil {
					all["question"] = q.String()
				}
				a := qa.GetA(ctx)
				if a != nil {
					all["answer"] = a.String()
					all["result"] = a.Run(ctx)
				}
				so.Data, _ = json.JsonParser().MarshalToString(&all)
			}
		}
	}

	out = so
	return
}

func (app *Application) init() {
	app.chapter = map[string]base2.Chapter{
		"map":   chapter1_map.New(),
		"slice": chapter2_slice.New(),
	}
}
