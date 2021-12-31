package search

import (
	"fmt"
	"graffito/experiment/search/lib"
)

func Run() {
	doc1 := []string{"人工", "智能", "成为", "互联网", "大会", "焦点"}
	doc2 := []string{"谷歌", "推出", "开源", "人工", "智能", "系统", "工具"}
	doc3 := []string{"互联网", "的", "未来", "在", "人工", "智能"}
	doc4 := []string{"谷歌", "开源", "机器", "学习", "工具"}

	docs := [][]string{doc1, doc2, doc3, doc4}

	tfIdf := lib.NewTfIdf(docs)
	fmt.Println("tf:", tfIdf.Tf(1, "谷歌"))
	fmt.Println("df:", tfIdf.Df("谷歌"))
	fmt.Println("tfidf:", tfIdf.TfIdf(1, "谷歌"))
}
