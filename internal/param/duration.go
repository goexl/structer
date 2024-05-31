package param

import (
	"reflect"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/structer/internal/internal/constant"
	"github.com/goexl/structer/internal/internal/variable"
)

type Duration struct {
	params *Copy
}

func NewDuration(params *Copy) *Duration {
	return &Duration{
		params: params,
	}
}

func (d *Duration) Protobuf(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	var seconds int64 = 0
	var nanos int32 = 0

	name := ft.String()
	if name == constant.DurationPtr && d.isProtobufType(tt) {
		duration := from.(*time.Duration)
		nano := duration.Nanoseconds()
		seconds = nano / 1000000000
		nanos = int32(nano % 1000000000)
	} else if name == constant.Duration && d.isProtobufType(tt) {
		duration := from.(time.Duration)
		nano := duration.Nanoseconds()
		seconds = nano / 1000000000
		nanos = int32(nano % 1000000000)
	}

	if 0 != nanos || 0 != seconds {
		secondsKey := gox.Ift(constant.Json == d.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
		nanosKey := gox.Ift(constant.Json == d.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)
		to = &map[string]any{
			secondsKey: seconds,
			nanosKey:   nanos,
		}
	} else {
		to = from
	}

	return
}

func (d *Duration) Internal(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	toName := tt.String()
	if toName != constant.Duration && toName != constant.DurationPtr {
		to = from
	} else if toName == constant.Duration && ft == variable.TypeMap {
		to = d.parseMap(from)
	} else if toName == constant.DurationPtr && ft == variable.TypeMap {
		toValue := d.parseMap(from)
		to = &toValue
	}

	return
}

func (*Duration) isProtobufType(tt reflect.Type) bool {
	return tt.String() == constant.PBDuration || tt.String() == constant.PBDurationPtr
}

func (d *Duration) parseMap(from any) time.Duration {
	data := from.(map[string]any)
	secondsKey := gox.Ift(constant.Json == d.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
	nanosKey := gox.Ift(constant.Json == d.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)

	var duration int64 = 0
	if value, ok := data[secondsKey]; ok {
		duration = duration + value.(int64)*time.Second.Nanoseconds()
	}
	if value, ok := data[nanosKey]; ok {
		duration = duration + value.(int64)
	}

	return time.Duration(duration)
}
