package structer

import (
	"github.com/mitchellh/mapstructure"
)

type _map struct {
	zero   bool
	squash bool
	tag    string

	from any
	to   any
}

func newMap() *_map {
	return &_map{
		zero:   true,
		squash: true,
		tag:    "json",
	}
}

func (m *_map) Convert() (err error) {
	config := mapstructure.DecoderConfig{
		ZeroFields: m.zero,
		Result:     m.to,
		Squash:     m.squash,
		TagName:    m.tag,
	}
	if decoder, ne := mapstructure.NewDecoder(&config); nil != ne {
		err = ne
	} else {
		err = decoder.Decode(m.from)
	}

	return
}

func (m *_map) From(from any) *_map {
	m.from = from

	return m
}

func (m *_map) To(to any) *_map {
	m.to = to

	return m
}

func (m *_map) DisableSquash() *_map {
	m.squash = false

	return m
}

func (m *_map) DisableZeroFields() *_map {
	m.zero = false

	return m
}

func (m *_map) Tag(tag string) *_map {
	m.tag = tag

	return m
}
