package base

import "github.com/2lovecode/graffito/pkg/json"

type IPayload interface {
	Serialization() (target []byte, err error)
	Deserialization(source []byte) (err error)
	From(source IPayload) (err error)
	To(target IPayload) (err error)
}

func Serialization(p IPayload) (target []byte, err error) {
	if p != nil {
		target, err = json.JsonParser().Marshal(p)
	}
	return
}

func Deserialization(source []byte, p IPayload) (err error) {
	if p != nil {
		err = json.JsonParser().Unmarshal(source, p)
	}
	return
}

func From(source IPayload, target IPayload) (err error) {
	sourceByte := make([]byte, 0)
	sourceByte, err = source.Serialization()
	if err == nil {
		err = target.Deserialization(sourceByte)
	}
	return
}

func To(source IPayload, target IPayload) (err error) {
	targetByte := make([]byte, 0)
	targetByte, err = source.Serialization()
	if err == nil {
		err = target.Deserialization(targetByte)
	}
	return
}
