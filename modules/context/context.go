package context

import (
	"path"

	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/modules/theme"
	macaron "gopkg.in/macaron.v1"
)

// Middleware adds our own context
func Middleware() macaron.Handler {
	renderers := Renderers{
		Admin: macaron.Renderer(macaron.RenderOptions{
			Directory: "templates",
			Funcs:     getTemplateFuncs(),
		}),
		Theme: macaron.Renderer(macaron.RenderOptions{
			Directory: path.Join(config.WorkingDirectory, theme.ThemeDir),
			Funcs:     getTemplateFuncs(),
		}),
	}

	return func(realCtx *macaron.Context) {
		ctx := Context{
			RealCtx:   realCtx,
			Renderers: renderers,
			Data:      realCtx.Data,
		}

		realCtx.Map(&ctx)
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
	Data      map[string]interface{}
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

// Status ...
func (c *Context) Status(code int) {
	c.RealCtx.Invoke(c.Renderers.Theme)
	c.RealCtx.Status(code)
}

// Redirect ...
func (c *Context) Redirect(url string, status ...int) {
	c.RealCtx.Redirect(url, status...)
}

// JSON ....
func (c *Context) JSON(code int, data interface{}) {
	c.RealCtx.Invoke(c.Renderers.Theme)
	c.RealCtx.JSON(code, data)
}
