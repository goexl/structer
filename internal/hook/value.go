package hook

import (
	"reflect"
)

type Value func(from reflect.Value, to reflect.Value) (any, error)
