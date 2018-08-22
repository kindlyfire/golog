package router

import (
	"github.com/go-macaron/session"
	"github.com/kindlyfire/golog/modules/context"
	"github.com/kindlyfire/golog/router/routes/index"

	"gopkg.in/macaron.v1"
)

// New creates a Macaron instance with all the middleware
func New() *macaron.Macaron {
	m := macaron.New()

	m.Use(macaron.Logger())
	m.Use(macaron.Static("public", macaron.StaticOptions{
		Prefix: "public",
	}))
	m.Use(session.Sessioner())
	m.Use(context.Middleware())

	return m
}

// Register registers all routes
func Register(m *macaron.Macaron) {
	m.Get("/", index.Index)
}
