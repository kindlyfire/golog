package posts

import (
	"strings"

	"github.com/go-macaron/binding"

	"github.com/jinzhu/gorm"
	"github.com/kindlyfire/golog/models"
	"github.com/kindlyfire/golog/modules/context"
)

// FindForm ...
type FindForm struct {
	Id      int
	Limit   int
	Offset  int
	Preload string
}

// FindJSON ...
type FindJSON struct {
	Success bool
	Data    struct {
		Posts []models.Post
		Count int
	}
}

// Find a post, or multiple posts
func Find(ctx *context.Context, db *gorm.DB, form FindForm, errs binding.Errors) {
	ret := FindJSON{Success: true}

	// Perform a count of all rows that match
	// Ignoring the limit and offset
	doMainFilters(db, form).Count(&ret.Data.Count)

	// We need to build the query again, but this time with
	db = doMainFilters(db, form)

	// Limit & offset
	if form.Limit != 0 {
		db = db.Limit(form.Limit)
	}
	db = db.Offset(form.Offset)

	// Preloads
	preloads := strings.Split(form.Preload, ",")
	for _, p := range preloads {
		if p == "Author" || p == "Tags" || p == "Categories" {
			db = db.Preload(p)
		}
	}

	posts := []models.Post{}
	db.Find(&posts)
	ret.Data.Posts = posts

	ctx.JSON(200, ret)
}

func doMainFilters(db *gorm.DB, form FindForm) *gorm.DB {
	db = db.Table("posts")
	if form.Id != 0 {
		db = db.Where("ID = ?", form.Id)
	}
	return db
}
