package rule

import "github.com/2lovecode/graffito/internal/app/base"

type Input struct {
	SourceCode string
}

func (in *Input) Serialization() (target []byte, err error) {
	return base.Serialization(in)
}

func (in *Input) Deserialization(source []byte) (err error) {
	return base.Deserialization(source, in)
}

func (in *Input) From(source base.IPayload) (err error) {
	return base.From(source, in)
}

func (in *Input) To(target base.IPayload) (err error) {
	return base.To(in, target)
}
