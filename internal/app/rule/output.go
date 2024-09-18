package rule

import "github.com/2lovecode/graffito/internal/app/base"

type Output struct {
	Data string
}

func (out *Output) Serialization() (target []byte, err error) {
	return base.Serialization(out)
}

func (out *Output) Deserialization(source []byte) (err error) {
	return base.Deserialization(source, out)
}

func (out *Output) From(source base.IPayload) (err error) {
	return base.From(source, out)
}

func (out *Output) To(target base.IPayload) (err error) {
	return base.To(out, target)
}
