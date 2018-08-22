package auth

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

// Login shows the login page
func Login(ctx *macaron.Context, sess session.Store) {
	sess.Set("user_id", 1)
	ctx.Redirect("/gl-admin")
}
