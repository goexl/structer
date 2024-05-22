package internal

import (
	"reflect"
	"time"
)

var (
	TypeDurationPtr = reflect.TypeOf((*time.Duration)(nil))
	TypeDuration    = TypeDurationPtr.Elem()
)
