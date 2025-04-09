package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	doc := struct {
		Id   string    `json:"id"`
		Text string    `json:"text"`
		Vec  []float32 `json:"vec"`
	}{
		Id:   "example",
		Text: "hello from united states",
		Vec:  []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	textFieldMapping := bleve.NewTextFieldMapping()
	// vectorFieldMapping := bleve.NewVectorFieldMapping()
	// vectorFieldMapping.Dims = 10
	// vectorFieldMapping.Similarity = "l2_norm" // euclidean distance

	bleveMapping := bleve.NewIndexMapping()
	bleveMapping.DefaultMapping.Dynamic = false
	bleveMapping.DefaultMapping.AddFieldMappingsAt("text", textFieldMapping)
	// bleveMapping.DefaultMapping.AddFieldMappingsAt("vec", vectorFieldMapping)

	// open a new index
	// mapping := bleve.NewIndexMapping()
	// mapping.StoreDynamic = true
	index, err := bleve.NewMemOnly(bleveMapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	// data := struct {
	// 	Name string `json:"name"`
	// 	Desc string `json:"desc"`
	// }{
	// 	Name: "清水好房",
	// 	Desc: "abc",
	// }

	// index some data
	index.Index(doc.Id, doc)

	// search for some text
	query := bleve.NewMatchQuery("hello")
	search := bleve.NewSearchRequest(query)
	// search.Highlight = bleve.NewHighlight()
	// search.Highlight.AddField("text")
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits)

	for _, hit := range searchResults.Hits {
		fmt.Println(hit.ID, hit.Score)
		fmt.Println(hit.Fields)
		fmt.Println(hit.Fields["_source"])
	}
}
