package models

import "github.com/jinzhu/gorm"

// Option is an option
type Option struct {
	gorm.Model
	Name  string `gorm:"type:varchar(64);unique_index"`
	Value string
}
