package posts

import (
	"github.com/jinzhu/gorm"
	"github.com/kindlyfire/golog/models"
	"github.com/kindlyfire/golog/modules/context"
)

type ListJson struct {
	Success bool
	Data    struct {
		Posts []models.Post
	}
}

// List all posts
func List(ctx *context.Context, db *gorm.DB) {
	ret := ListJson{}
	ret.Success = true

	posts := []models.Post{}
	db.Preload("Author").Find(&posts)
	ret.Data.Posts = posts

	ctx.JSON(200, ret)
}
