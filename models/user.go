package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User is a user
type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20)"`
	Name     string `gorm:"type:varchar(32)"`
	Email    string `gorm:"type:varchar(32)"`
	Password string `gorm:"type:varchar(60)" json:"-"`

	Posts []Post `gorm:"foreignkey:author_id"`
}

// SetPassword sets the password for a user
func (u *User) SetPassword(pass string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 11)

	if err == nil {
		u.Password = fmt.Sprintf("%s", bytes)
	}

	return err
}

// CheckPassword checks if a password matches to the password hash
func (u User) CheckPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}
