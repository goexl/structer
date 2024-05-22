package core

import (
	"github.com/goexl/structer/internal/param"
	"github.com/mitchellh/mapstructure"
)

type Clone struct {
	params *param.Clone
}

func NewClone(params *param.Clone) *Clone {
	return &Clone{
		params: params,
	}
}

func (c *Clone) Apply() (err error) {
	config := new(mapstructure.DecoderConfig)
	config.ZeroFields = c.params.Zero
	config.Result = c.params.To
	config.Squash = c.params.Squash
	config.TagName = c.params.Tag
	config.ErrorUnused = c.params.Unused
	config.ErrorUnset = c.params.Unset
	config.WeaklyTypedInput = c.params.Weakly
	if nil != c.params.Hook {
		config.DecodeHook = c.params.Hook
	}

	if decoder, ne := mapstructure.NewDecoder(config); nil != ne {
		err = ne
	} else {
		err = decoder.Decode(c.params.From)
	}

	return
}
