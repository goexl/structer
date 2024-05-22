package internal

import (
	"reflect"
	"time"
)

var (
	TypeTimePtr = reflect.TypeOf((*time.Time)(nil))
	TypeTime    = TypeTimePtr.Elem()
)
