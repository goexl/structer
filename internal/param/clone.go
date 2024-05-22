package param

import (
	"github.com/goexl/structer/internal/constant"
)

type Clone struct {
	Zero   bool
	Squash bool
	Tag    string
	Unused bool
	Unset  bool

	From any
	To   any

	Hook   any
	Weakly bool
}

func NewClone() *Clone {
	return &Clone{
		Zero:   true,
		Squash: true,
		Tag:    constant.Json,
		Weakly: true,
	}
}
