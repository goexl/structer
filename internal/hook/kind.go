package hook

import (
	"reflect"
)

type Kind func(reflect.Kind, reflect.Kind, any) (any, error)
