package common

import (
	"github.com/jinzhu/gorm"
)

// DB is pointer of active database pq
var DB *gorm.DB

// Product sample template of how gorm tables are
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// SetDB pointer from active database
func SetDB(db *gorm.DB) {
	DB = db
	if res :=DB.HasTable(&Product{}); res == false {
		DB.AutoMigrate(&Product{})
	}
}