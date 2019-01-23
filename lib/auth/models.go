package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username  string
	Password string
	Email string
}

type JWT struct {
	User User
	Age int
	jwt.StandardClaims
}

// example usage of this two functions

// pass := SetPassword(password)
// bo := ValidPassword(pass, "as")
// fmt.Println(pass, bo)

func SetPassword(pass string) string {
	return hashAndSalt([]byte(pass))
}

func ValidPassword(hash string, pass string) bool {
	return comparePasswords(hash, []byte(pass))
}

func SetDB(db *gorm.DB) {
	DB = db
	if res :=DB.HasTable(&User{}); res == false {
		DB.AutoMigrate(&User{})
	}
	if res :=DB.HasTable(&JWT{}); res == false {
		DB.AutoMigrate(&JWT{})
	}
}
