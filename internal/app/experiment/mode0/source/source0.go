package source

import (
	"github.com/2lovecode/graffito/internal/app/experiment/mode0"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
)

type S0Resp struct {
	Type string
	Name string
}
type S0 struct {
	Resp []S0Resp
}

func NewS0() S0 {

	s := S0{}
	return s
}

func (s S0) Output() mode0.OutData {

	//mock data
	s.Resp = []S0Resp{
		{
			Type: "big",
			Name: "type_big",
		},
		{
			Type: "small",
			Name: "type_small",
		},
	}
	//

	output := mode0.OutData{
		CommonData: nil,
		Data:       make([]handler.DataItem, len(s.Resp)),
	}
	for idx, r := range s.Resp {
		output.Data[idx].Type = r.Type
		output.Data[idx].Name = r.Name
	}

	return output
}
