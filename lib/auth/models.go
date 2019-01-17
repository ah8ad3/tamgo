package auth

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	username  string
	password string
	email string
}

func (user *User) SetPassword(pass string) string {
	return hashAndSalt([]byte(pass))
}

func (user *User) ValidPassword(pass string, check string) bool {
	return comparePasswords(pass, []byte(check))
}

func SetDB(db *gorm.DB) {
	DB = db
}
