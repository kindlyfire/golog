package context

import (
	macaron "gopkg.in/macaron.v1"
)

// Middleware adds our own context
func Middleware() macaron.Handler {
	return func(realCtx *macaron.Context) {
		ctx := Context{
			RealCtx: *realCtx,
		}

		realCtx.Map(ctx)
	}
}

// Context is our own context
type Context struct {
	RealCtx macaron.Context
}

// Render renders files our way
func (c *Context) Render(code int, tpl string) {
	c.RealCtx.HTML(code, tpl)
}
