package param

import (
	"github.com/goexl/structer/internal/constant"
	"github.com/mitchellh/mapstructure"
)

type Clone struct {
	Zero     bool
	Squash   bool
	Tag      string
	Unused   bool
	Unset    bool
	Untagged bool

	From any
	To   any

	Hooks  []mapstructure.DecodeHookFunc
	Weakly bool
}

func NewClone() *Clone {
	return &Clone{
		Zero:     true,
		Squash:   true,
		Tag:      constant.Json,
		Untagged: true,

		Hooks:  make([]mapstructure.DecodeHookFunc, 0),
		Weakly: true,
	}
}
