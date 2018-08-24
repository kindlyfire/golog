package models

import "github.com/jinzhu/gorm"

// Option is an option
type Post struct {
	gorm.Model

	Author   User
	AuthorID uint

	Title         string `gorm:"type:varchar(255)"`
	Excerpt       string `gorm:"type:varchar(255)"`
	Content       string
	Status        string `gorm:"type:varchar(16)"`
	CommentStatus string `gorm:"type:varchar(16)"`
	CommentCount  uint
}
