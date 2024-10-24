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
		time.Protobuf, // 从 time.Time 转换到 timepb.Time
		time.Internal, // 从 timepb.Time 转换到 time.Time

		duration.Protobuf, // 从 time.Duration 转换到 durationpb.Duration
		duration.Internal, // 从 durationpb.Duration 转换到 time.Duration

		bytes.Protobuf,
		bytes.Internal,
	}

	return
}
