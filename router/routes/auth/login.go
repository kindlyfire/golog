package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/kindlyfire/golog/models"
	"github.com/kindlyfire/golog/modules/context"

	"github.com/go-macaron/session"
)

// Login shows the login page
func Login(ctx *context.Context, sess session.Store) {
	if ctx.Data["LoggedIn"].(bool) {
		ctx.Redirect("/")
		return
	}

	ctx.AdminHTML(200, "auth/login")
}

// LoginPostForm is the form for LoginPost
type LoginPostForm struct {
	UID      string `form:"uid" binding:"Required" forwardAs:"ForwardedUID" err:"Username or email required"`
	Password string `form:"password" binding:"Required" err:"Password is required"`
}

// Example of validation
// func (f LoginPostForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
// 	if f.UID == "a" {
// 		errs = append(errs, binding.Error{
// 			FieldNames:     []string{"UID"},
// 			Classification: "Custom",
// 			Message:        "Test message !",
// 		})
// 	}
// 	return errs
// }

// LoginPost ...
func LoginPost(ctx *context.Context, db *gorm.DB, form LoginPostForm, flash *session.Flash, sess session.Store) {
	user := models.User{}
	err := db.Where("username = ? OR email = ?", form.UID, form.UID).First(&user).Error

	if err != nil || !user.CheckPassword(form.Password) {
		// Flash an error
		flash.Error("Invalid username or password.")

		// Forward some data
		sess.Set("ForwardedUID", form.UID)
		sess.Set("Forwards", "ForwardedUID")

		// Redirect
		ctx.Redirect("/gl-auth/login")
		return
	}

	sess.Set("user_id", user.ID)
	ctx.Redirect("/gl-admin")
}
