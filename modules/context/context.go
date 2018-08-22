package context

import (
	"github.com/kindlyfire/golog/modules/theme"
	macaron "gopkg.in/macaron.v1"
)

// Middleware adds our own context
func Middleware() macaron.Handler {
	renderers := Renderers{
		Admin: macaron.Renderer(macaron.RenderOptions{
			Directory: "templates",
		}),
		Theme: theme.GetRenderer(),
	}

	return func(realCtx *macaron.Context) {
		ctx := Context{
			RealCtx:   realCtx,
			Renderers: renderers,
		}

		realCtx.Map(ctx)
	}
}

// Renderers are the handlers we call to create the ctx.HTML method
type Renderers struct {
	Admin macaron.Handler
	Theme macaron.Handler
}

// Context is our own context
type Context struct {
	RealCtx   *macaron.Context
	Renderers Renderers
}

// ThemeHTML renders theme template
func (c *Context) ThemeHTML(code int, tpl string) {
	c.RealCtx.Invoke(c.Renderers.Theme)
	c.RealCtx.HTML(code, tpl)
}

// AdminHTML renders admin template
func (c *Context) AdminHTML(code int, tpl string) {
	c.RealCtx.Invoke(c.Renderers.Admin)
	c.RealCtx.HTML(code, tpl)
}
