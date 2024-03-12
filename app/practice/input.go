package practice

import "graffito/app/base"

type Input struct {
	Chapter string // 章节
	Number  int    // 题目序号
	Part    string // q-问题，a-答案，all-问题+答案+执行结果
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
