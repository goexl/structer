package structer

type converterBuilder struct {
	params *converterParams
}

func newConverterBuilder() *converterBuilder {
	return &converterBuilder{
		params: newConverterParams(),
	}
}

func (cb *converterBuilder) From(from any) *converterBuilder {
	cb.params.from = from

	return cb
}

func (cb *converterBuilder) To(to any) *converterBuilder {
	cb.params.to = to

	return cb
}

func (cb *converterBuilder) DisableSquash() *converterBuilder {
	cb.params.squash = false

	return cb
}

func (cb *converterBuilder) DisableZeroFields() *converterBuilder {
	cb.params.zero = false

	return cb
}

func (cb *converterBuilder) Tag(tag string) *converterBuilder {
	cb.params.tag = tag

	return cb
}

func (cb *converterBuilder) Build() *converter {
	return newConverter(cb.params)
}
