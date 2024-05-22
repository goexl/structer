package structer

import (
	"github.com/goexl/structer/internal/builder"
)

func Clone() *builder.Clone {
	return builder.NewClone()
}
