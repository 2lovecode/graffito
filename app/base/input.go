package base

//type Input struct {
//	payload Payload
//}
//
//func (in *Input) Wrap(payload Payload) {
//	in.payload = payload
//	return
//}
//
//func (in *Input) Unwrap(receiver interface{}) (err error) {
//	if in.payload != nil {
//		bt := make([]byte, 0)
//		bt, err = in.payload.Serialization()
//		if err == nil && len(bt) > 0 {
//			err = json.JsonParser().Unmarshal(bt, receiver)
//		}
//	}
//	return
//}

type Input interface {
	IPayload
}
