package param

import (
	"github.com/goexl/structer/internal/internal/callback"
	"github.com/goexl/structer/internal/internal/constant"
	"github.com/mitchellh/mapstructure"
)

type Copy struct {
	Zero     bool
	Squash   bool
	Tag      string
	Unused   bool
	Unset    bool
	Untagged bool
	Mapper   callback.Mapper

	From any
	To   any

	Hooks  []mapstructure.DecodeHookFunc
	Weakly bool
}

func NewCopy() (copy *Copy) {
	copy = new(Copy)
	copy.Squash = true
	copy.Tag = constant.Json
	copy.Weakly = true

	time := NewTime(copy)
	duration := NewDuration(copy)
	bytes := NewBytes(copy)
	copy.Hooks = []mapstructure.DecodeHookFunc{
		time.Protobuf,     // 从time.Time转换到timepb.Time
		time.Internal,     // 从timepb.Time转换到time.Time
		duration.Protobuf, // 从time.Duration转换到durationpb.Duration
		duration.Internal, // 从durationpb.Duration转换到time.Duration

		bytes.Protobuf,
		bytes.Internal,
	}

	return
}
