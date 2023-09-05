package base

//type Output struct {
//	payload Payload
//}
//
//func (out *Output) Wrap(payload Payload) {
//	out.payload = payload
//	return
//}
//
//func (out *Output) Unwrap(receiver interface{}) (err error) {
//	if out.payload != nil {
//		bt := make([]byte, 0)
//		bt, err = out.payload.Serialization()
//		if err == nil && len(bt) > 0 {
//			err = json.JsonParser().Unmarshal(bt, receiver)
//		}
//	}
//	return
//}

type Output interface {
	IPayload
}
