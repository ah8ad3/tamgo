package auth

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username  string
	Password string
	Email string
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

func CreateAuthModel() {
	//DB.Create(&User{Username: "ahmad", Password: "ahmad", Email: "ahmad@gmail.com"})
	if res :=DB.HasTable(&User{}); res == false {
		fmt.Println(&User{})
		DB.AutoMigrate(&User{})
	}
}
