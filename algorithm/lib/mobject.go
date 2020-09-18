package lib

import (
	"errors"
	"reflect"
)

type MInterface interface {
	Compare(his MInterface) int8
}

type MInteger int

func (mine MInteger) Compare(his MInteger) int8 {
	if mine == his {
		return 0
	} else if mine > his {
		return 1
	}
	return -1
}

type MObject struct {
	Value MInterface
}

func (mObj MObject) Compare(hObj MObject) int8 {
	return mObj.Value.Compare(hObj.Value)
}

func (mObj *MObject) Get(receiver interface{}) error {
	receiverV := reflect.ValueOf(receiver)
	if receiverV.Kind() != reflect.Ptr {
		return errors.New("non-pointer")
	}

	receiverV = receiverV.Elem()

	receiverT := receiverV.Type()

	realV := reflect.ValueOf(mObj.Value)

	if !realV.Type().ConvertibleTo(receiverT) {
		return errors.New("can not convert")
	}

	receiverV.Set(realV.Convert(receiverT))

	return nil
}