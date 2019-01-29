package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// DB is a pointer of active database pq
var DB *gorm.DB

// User model for save in database
type User struct {
	gorm.Model
	Username  string
	Password string
	Email string
}

// JWT model for save in database
type JWT struct {
	User User
	Age int
	jwt.StandardClaims
}

// SetPassword example usage of this two functions
// pass := SetPassword(password)
func SetPassword(pass string) string {
	return hashAndSalt([]byte(pass))
}

// ValidPassword bo := ValidPassword(pass, "as")
// fmt.Println(pass, bo)
func ValidPassword(hash string, pass string) bool {
	return comparePasswords(hash, []byte(pass))
}

// SetDB this is how you can set DB pointer to active database
func SetDB(db *gorm.DB) {
	DB = db
	if res :=DB.HasTable(&User{}); res == false {
		DB.AutoMigrate(&User{})
	}
	if res :=DB.HasTable(&JWT{}); res == false {
		DB.AutoMigrate(&JWT{})
	}
}
