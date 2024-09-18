package builder

import (
	"github.com/goexl/structer/internal/core"
	"github.com/goexl/structer/internal/hook"
	"github.com/goexl/structer/internal/internal/callback"
	"github.com/goexl/structer/internal/internal/constant"
	"github.com/goexl/structer/internal/param"
)

type Copy struct {
	params *param.Copy
}

func NewCopy() *Copy {
	return &Copy{
		params: param.NewCopy(),
	}
}

func (c *Copy) From(from any) (copy *Copy) {
	c.params.From = from
	copy = c

	return
}

func (c *Copy) To(to any) (copy *Copy) {
	c.params.To = to
	copy = c

	return
}

func (c *Copy) DisableSquash() (copy *Copy) {
	c.params.Squash = false
	copy = c

	return
}

func (c *Copy) Zero() (copy *Copy) {
	c.params.Zero = true
	copy = c

	return
}

func (c *Copy) DisableWeakly() (copy *Copy) {
	c.params.Weakly = false
	copy = c

	return
}

func (c *Copy) DisableUntagged() (copy *Copy) {
	c.params.Untagged = false
	copy = c

	return
}

func (c *Copy) ErrorOnUnused() (copy *Copy) {
	c.params.Unused = true
	copy = c

	return
}

func (c *Copy) ErrorOnUnset() (copy *Copy) {
	c.params.Unset = true
	copy = c

	return
}

func (c *Copy) Mapper(mapper callback.Mapper) (copy *Copy) {
	c.params.Mapper = mapper
	copy = c

	return
}

func (c *Copy) Type(required hook.Type, others ...hook.Type) (copy *Copy) {
	c.params.Hooks = append(c.params.Hooks, required)
	for _, other := range others {
		c.params.Hooks = append(c.params.Hooks, other)
	}
	c.params.Weakly = true
	copy = c

	return
}

func (c *Copy) Kind(required hook.Kind, others ...hook.Kind) (copy *Copy) {
	c.params.Hooks = append(c.params.Hooks, required)
	for _, other := range others {
		c.params.Hooks = append(c.params.Hooks, other)
	}
	c.params.Weakly = true
	copy = c

	return
}

func (c *Copy) Value(required hook.Value, others ...hook.Value) (copy *Copy) {
	c.params.Hooks = append(c.params.Hooks, required)
	for _, other := range others {
		c.params.Hooks = append(c.params.Hooks, other)
	}
	c.params.Weakly = true
	copy = c

	return
}

func (c *Copy) Tag(tag string) (copy *Copy) {
	c.params.Tag = tag
	copy = c

	return
}

func (c *Copy) Json() *Copy {
	return c.Tag(constant.Json)
}

func (c *Copy) Struct() (copy *Copy) {
	return c.Tag(constant.DefaultTag)
}

func (c *Copy) Build() *core.Copy {
	return core.NewCopy(c.params)
}
