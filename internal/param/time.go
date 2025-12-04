package param

import (
	"reflect"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/structer/internal/internal/constant"
	"github.com/goexl/structer/internal/internal/variable"
)

type Time struct {
	params *Copy
}

func NewTime(params *Copy) *Time {
	return &Time{
		params: params,
	}
}

func (t *Time) Protobuf(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	var seconds int64 = 0
	var nanos int32 = 0

	name := ft.String()
	if name == constant.TimePtr && t.isProtobufType(tt) {
		ts := from.(*time.Time)
		nano := ts.UnixNano()
		seconds = nano / 1000000000
		nanos = int32(nano % 1000000000)
	} else if name == constant.Time && t.isProtobufType(tt) {
		ts := from.(time.Time)
		nano := ts.UnixNano()
		seconds = nano / 1000000000
		nanos = int32(nano % 1000000000)
	}

	if 0 != nanos || 0 != seconds {
		secondsKey := gox.Ift(constant.Json == t.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
		nanosKey := gox.Ift(constant.Json == t.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)
		to = &map[string]any{
			secondsKey: seconds,
			nanosKey:   nanos,
		}
	} else {
		to = from
	}

	return
}

func (t *Time) Internal(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	name := tt.String()
	if name != constant.TimePtr && name != constant.Time {
		to = from
	} else if name == constant.Time && ft == variable.TypeMap {
		data := from.(map[string]any)
		secondsKey := gox.Ift(constant.Json == t.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
		nanosKey := gox.Ift(constant.Json == t.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)

		var timestamp int64 = 0
		if value, ok := data[secondsKey]; ok {
			timestamp = timestamp + value.(int64)*time.Second.Nanoseconds()
		}
		if value, ok := data[nanosKey]; ok {
			timestamp = timestamp + value.(int64)
		}
		to = time.Unix(0, timestamp)
	}

	return
}

func (*Time) isProtobufType(tt reflect.Type) bool {
	return tt.String() == constant.PBTime || tt.String() == constant.PBTimePtr
}
