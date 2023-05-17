package structer

type converterParams struct {
	zero   bool
	squash bool
	tag    string

	from any
	to   any
}

func newConverterParams() *converterParams {
	return &converterParams{
		zero:   true,
		squash: true,
		tag:    "json",
	}
}
