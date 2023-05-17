package structer

var _ = Converter

func Converter() *converterBuilder {
	return newConverterBuilder()
}
