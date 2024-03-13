package search

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/base"
	"github.com/2lovecode/graffito/internal/app/search/pkg"
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
	doc1 := []string{"人工", "智能", "成为", "互联网", "大会", "焦点"}
	doc2 := []string{"谷歌", "推出", "开源", "人工", "智能", "系统", "工具"}
	doc3 := []string{"互联网", "的", "未来", "在", "人工", "智能"}
	doc4 := []string{"谷歌", "开源", "机器", "学习", "工具"}

	docs := [][]string{doc1, doc2, doc3, doc4}

	tfIdf := pkg.NewTfIdf(docs)
	fmt.Println("tf:", tfIdf.Tf(1, "谷歌"))
	fmt.Println("df:", tfIdf.Df("谷歌"))
	fmt.Println("tfidf:", tfIdf.TfIdf(1, "谷歌"))

	out = so
	return
}
