package hook

import (
	"reflect"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/structer/internal/constant"
	"github.com/goexl/structer/internal/hook/internal"
	"github.com/goexl/structer/internal/param"
)

type Time struct {
	params *param.Copy
}

func NewTime(params *param.Copy) *Time {
	return &Time{
		params: params,
	}
}

func (t *Time) TimeToPB(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	var seconds int64 = 0
	var nanos int32 = 0
	if ft == internal.TypeTimePtr && tt == internal.TypeMap {
		ts := from.(*time.Time)
		nano := ts.UnixNano()
		seconds = nano / 1000000000
		nanos = int32(nano % 1000000000)
	}

	if 0 != nanos {
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

func (t *Time) DurationToPB(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	var seconds float64 = 0
	var nanos int64 = 0
	if ft == internal.TypeDuration && tt == internal.TypeMap {
		duration := from.(time.Duration)
		seconds = duration.Seconds()
		nanos = duration.Nanoseconds()
	} else if ft == internal.TypeDurationPtr && tt == internal.TypeMap {
		duration := from.(*time.Duration)
		seconds = duration.Seconds()
		nanos = duration.Nanoseconds()
	}

	if 0 != nanos {
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

func (t *Time) PBToDuration(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	if tt != internal.TypeDuration {
		to = from
	} else if tt == internal.TypeDuration && ft == internal.TypeMap {
		data := from.(map[string]any)
		secondsKey := gox.Ift(constant.Json == t.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
		nanosKey := gox.Ift(constant.Json == t.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)

		var duration int64 = 0
		if value, ok := data[secondsKey]; ok {
			duration = duration + value.(int64)*time.Second.Nanoseconds()
		}
		if value, ok := data[nanosKey]; ok {
			duration = duration + value.(int64)
		}
		to = time.Duration(duration)
	}

	return
}

func (t *Time) PBToTime(ft reflect.Type, tt reflect.Type, from any) (to any, err error) {
	if tt != internal.TypeTimePtr || tt != internal.TypeTime {
		to = from
	} else if tt == internal.TypeTime && ft == internal.TypeMap {
		data := from.(map[string]any)
		secondsKey := gox.Ift(constant.Json == t.params.Tag, constant.KeySeconds, constant.KeySecondsUpper)
		nanosKey := gox.Ift(constant.Json == t.params.Tag, constant.KeyNanos, constant.KeyNanosUpper)

		var ts int64 = 0
		if value, ok := data[secondsKey]; ok {
			ts = ts + value.(int64)*time.Second.Nanoseconds()
		}
		if value, ok := data[nanosKey]; ok {
			ts = ts + value.(int64)
		}
		to = time.Unix(0, ts)
	}

	return
}
