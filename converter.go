package structer

var _ = New

type converter struct{}

func New() *converter {
	return new(converter)
}

func (c *converter) Map() *_map {
	return newMap()
}
