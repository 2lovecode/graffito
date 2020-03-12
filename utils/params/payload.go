package params

import "reflect"

//载荷
type IPayload interface {
	Set(key string, value interface{})
	Get(key string, def interface{}) (value interface{})
	Has(key string) bool
}

//实现IPayload
type Payload struct {
	kvMap map[string]interface{}
}

func NewPayload () *Payload {
	return &Payload{
		kvMap:make(map[string]interface{}),
	}
}

func (mine *Payload) Set(key string, value interface{}) {
	mine.kvMap[key] = value
}

//根据传入的默认值的类型判断
func (mine *Payload) Get(key string, def interface{}) (value interface{}) {
	v, ok := mine.kvMap[key]
	if ok && reflect.TypeOf(def) == reflect.TypeOf(v) {
		return v
	}
	return def
}

func (mine *Payload) Has(key string) bool {
	_, ok := mine.kvMap[key]
	return ok
}
