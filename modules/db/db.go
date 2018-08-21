package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// ...
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kindlyfire/golog/modules/config"
)

var (
	// Db is the database
	Db *gorm.DB
)

// Connect connects the database
func Connect() {
	var err error

	Db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Db.User,
		config.Db.Password,
		config.Db.Addr,
		config.Db.Name,
	))

	if err != nil {
		panic(err)
	}
}
