package core

import (
	"github.com/goexl/structer/internal/param"
	"github.com/mitchellh/mapstructure"
)

type Copy struct {
	params *param.Copy
}

func NewCopy(params *param.Copy) *Copy {
	return &Copy{
		params: params,
	}
}

func (c *Copy) Apply() (err error) {
	config := new(mapstructure.DecoderConfig)
	config.ZeroFields = c.params.Zero
	config.Result = c.params.To
	config.Squash = c.params.Squash
	config.TagName = c.params.Tag
	config.ErrorUnused = c.params.Unused
	config.ErrorUnset = c.params.Unset
	config.WeaklyTypedInput = c.params.Weakly
	config.IgnoreUntaggedFields = c.params.Untagged
	if 0 != len(c.params.Hooks) {
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(c.params.Hooks...)
	}

	if decoder, ne := mapstructure.NewDecoder(config); nil != ne {
		err = ne
	} else {
		err = decoder.Decode(c.params.From)
	}

	return
}
