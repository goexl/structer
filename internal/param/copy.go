package param

import (
	"github.com/goexl/structer/internal/constant"
	"github.com/mitchellh/mapstructure"
)

type Copy struct {
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

func NewCopy() (copy *Copy) {
	copy = new(Copy)
	copy.Squash = true
	copy.Tag = constant.Json
	copy.Untagged = true
	copy.Weakly = true

	time := NewTime(copy)
	copy.Hooks = []mapstructure.DecodeHookFunc{
		time.PBToTime,
		time.PBToDuration,
		time.TimeToPB,
		time.DurationToPB,
	}

	return
}
