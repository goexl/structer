package param

import (
	"reflect"

	"github.com/goexl/gox"
	"github.com/goexl/structer/internal/internal/constant"
)

type Bytes struct {
	params *Copy
}

func NewBytes(params *Copy) *Bytes {
	return &Bytes{
		params: params,
	}
}

func (b *Bytes) Protobuf(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	name := ft.String()
	if name == constant.BytesPtr && reflect.String == tt.Kind() {
		value := from.(*gox.Bytes)
		to = value.String()
	} else if name == constant.Bytes && reflect.String == tt.Kind() {
		value := from.(gox.Bytes)
		to = value.String()
	} else {
		to = from
	}

	return
}

func (b *Bytes) Internal(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	name := tt.String()
	if name != constant.BytesPtr && name != constant.Bytes {
		to = from
	} else if name == constant.Bytes && reflect.String == ft.Kind() {
		data := from.(string)
		to, err = gox.ParseBytes(data)
	} else if name == constant.BytesPtr && reflect.String == ft.Kind() {
		data := from.(string)
		if parsed, pe := gox.ParseBytes(data); nil != pe {
			err = pe
		} else {
			to = &parsed
		}
	}

	return
}
