package router

import (
	"path"

	"github.com/go-macaron/session"
	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/modules/context"
	"github.com/kindlyfire/golog/modules/db"
	"github.com/kindlyfire/golog/modules/theme"
	"github.com/kindlyfire/golog/router/middleware"
	"github.com/kindlyfire/golog/router/routes/admin"
	"github.com/kindlyfire/golog/router/routes/auth"
	"github.com/kindlyfire/golog/router/routes/pages"

	"gopkg.in/macaron.v1"
)

// New creates a Macaron instance with all the middleware
func New() *macaron.Macaron {
	m := macaron.New()

	// Logger
	m.Use(macaron.Logger())

	// Golog public files
	m.Use(macaron.Static("public", macaron.StaticOptions{
		Prefix: "/_/gl-data",
	}))

	// Theme public files
	m.Use(macaron.Static(path.Join(config.WorkingDirectory, theme.ThemeDir, "public"), macaron.StaticOptions{
		Prefix: "/_/th-data",
	}))

	// Session support
	m.Use(session.Sessioner())

	// Build and map our context to the request
	m.Use(context.Middleware())

	// Fetch the user object if the user is logged in
	// And set ctx.Data["LoggedIn"] to true/false
	m.Use(middleware.FetchUser)

	// Forward data from a form that has had an error
	// To make it available
	m.Use(middleware.BindErrForwarder)

	// Map our database
	m.Map(db.Db)

	return m
}

// Register registers all routes
func Register(m *macaron.Macaron) {
	// Auth routes routes
	m.Group("/gl-auth", func() {
		m.Get("/login", auth.Login)
		m.Post("/login", middleware.Bind(auth.LoginPostForm{}), auth.LoginPost)
	})

	// Administration panel routes
	m.Group("/gl-admin", func() {
		m.Get("/", admin.Index)
	}, middleware.RequireUser)

	// Register index and catch-all page renderer
	m.Any("/", pages.Index)
	m.Any("/*", pages.Page)
}
