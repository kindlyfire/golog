package middleware

import (
	"github.com/go-macaron/session"
	"github.com/jinzhu/gorm"
	"github.com/kindlyfire/golog/models"
	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/modules/context"
)

// FetchUser fetches the user object if the user is logged in
func FetchUser(ctx *context.Context, sess session.Store, db *gorm.DB) {
	userID := sess.Get("user_id")
	var user models.User

	if userID != nil {
		user = models.User{}
		err := db.First(&user, userID).Error

		ctx.Data["LoggedIn"] = err == nil
	} else {
		ctx.Data["LoggedIn"] = false
	}

	ctx.Data["User"] = user
}

// RequireUser redirects to the homepage if the user is not logged in
func RequireUser(ctx *context.Context, flash *session.Flash) {
	if !ctx.Data["LoggedIn"].(bool) {
		ctx.Redirect(config.BaseUrl + "/gl-auth/login")
		return
	}
}
