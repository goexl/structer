package builder

import (
	"github.com/goexl/structer/internal/constant"
	"github.com/goexl/structer/internal/core"
	"github.com/goexl/structer/internal/hook"
	"github.com/goexl/structer/internal/param"
)

type Clone struct {
	params *param.Clone
}

func NewClone() *Clone {
	return &Clone{
		params: param.NewClone(),
	}
}

func (c *Clone) From(from any) (clone *Clone) {
	c.params.From = from
	clone = c

	return
}

func (c *Clone) To(to any) (clone *Clone) {
	c.params.To = to
	clone = c

	return
}

func (c *Clone) DisableSquash() (clone *Clone) {
	c.params.Squash = false
	clone = c

	return
}

func (c *Clone) DisableZero() (clone *Clone) {
	c.params.Zero = false
	clone = c

	return
}

func (c *Clone) DisableWeakly() (clone *Clone) {
	c.params.Weakly = false
	clone = c

	return
}

func (c *Clone) DisableUntagged() (clone *Clone) {
	c.params.Untagged = false
	clone = c

	return
}

func (c *Clone) ErrorOnUnused() (clone *Clone) {
	c.params.Unused = true
	clone = c

	return
}

func (c *Clone) ErrorOnUnset() (clone *Clone) {
	c.params.Unset = true
	clone = c

	return
}

func (c *Clone) Type(hook hook.Type) (clone *Clone) {
	c.params.Hooks = append(c.params.Hooks, hook)
	clone = c

	return
}

func (c *Clone) Kind(hook hook.Kind) (clone *Clone) {
	c.params.Hooks = append(c.params.Hooks, hook)
	clone = c

	return
}

func (c *Clone) Value(hook hook.Value) (clone *Clone) {
	c.params.Hooks = append(c.params.Hooks, hook)
	clone = c

	return
}

func (c *Clone) Tag(tag string) (clone *Clone) {
	c.params.Tag = tag
	clone = c

	return
}

func (c *Clone) Json() *Clone {
	return c.Tag(constant.Json)
}

func (c *Clone) Str() (clone *Clone) {
	return c.Tag(constant.DefaultTag)
}

func (c *Clone) Build() *core.Clone {
	return core.NewClone(c.params)
}
