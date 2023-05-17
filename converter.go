package structer

import (
	"github.com/mitchellh/mapstructure"
)

type converter struct {
	params *converterParams
}

func newConverter(params *converterParams) *converter {
	return &converter{
		params: params,
	}
}

func (c *converter) Convert() (err error) {
	config := mapstructure.DecoderConfig{
		ZeroFields: c.params.zero,
		Result:     c.params.to,
		Squash:     c.params.squash,
		TagName:    c.params.tag,
	}
	if decoder, ne := mapstructure.NewDecoder(&config); nil != ne {
		err = ne
	} else {
		err = decoder.Decode(c.params.from)
	}

	return
}
