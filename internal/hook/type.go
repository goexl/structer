package hook

import (
	"reflect"
)

type Type func(reflect.Type, reflect.Type, any) (any, error)
